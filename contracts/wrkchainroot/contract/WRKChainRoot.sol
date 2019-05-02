pragma solidity >=0.4.25;

import "./SafeMath.sol";

contract WRKChainRoot {

    // because it'd be daft not to
    using SafeMath for uint256;

    /**
    * ------
    * EVENTS
    * ------
    */

    // RecordHeader - called when block hashes are recorded. Allows querying the chain for hashes
    // without requiring to store that actual data.
    event RecordHeader(
        uint256  indexed _chainId,
        uint256  indexed _height,
        bytes32  indexed _hash,
        bytes32 _parentHash,
        bytes32 _receiptRoot,
        bytes32 _txRoot,
        bytes32 _stateRoot,
        address _blockSigner
    );

    event RegisterWrkChain(
        uint256  indexed _chainId,
        bytes32  _genesisHash
    );

    event LogFallbackFunctionCalled(
        address indexed _from,
        uint256 _amount
    );

    event WRKChainDepositRefund(
        uint256  indexed _chainId,
        address  indexed _owner,
        uint256 _amount
    );

    event AuthoriseNewAddress(
        uint256  indexed _chainId,
        address  indexed _authorisedBy,
        address _address
    );

    event RemoveAuthorisedAddress(
        uint256  indexed _chainId,
        address  indexed _removedBy,
        address _address
    );

    event WRKChainOwnerChanged(
        uint256  indexed _chainId,
        address _old,
        address _new
    );

    event WRKChainDepositPaid(
        uint256 indexed _chainId,
        address indexed _owner,
        uint256 _amount
    );

    /**
    * ----------
    * STRUCTURES
    * ----------
    */

    /**
    * @dev Struct to hold an individual WRKChain's high level data
    */
    struct Wrkchain {
        uint256 lastBlock; // Last block num submitted
        bytes32 genesisHash; // Hashed genesis block
        bool isWrkchain; // used in require statements to check if WRKChain exists
        mapping(address => bool) authAddresses; // Current wallet addresses allowed to write hashes
        address owner; // WRKChain owner address (the address that registered the WRKChain)
                       // This is the address that pays the deposit, and receives the refunded UND
        uint256 numBlocksSubmitted; // Counter for number of blocks submitted
    }

    /**
    * ---------
    * VARIABLES
    * ---------
    */

    /*
    * UND Deposit amount required for registering a WRKChain in Wei.
    * Set during deployment
    */
    uint256 private WRKCHAIN_REG_DEPOSIT;

    /*
    * Minimum number of blocks a WRKChain needs to submit before getting deposit refunded.
    * This is to ensure genuine WRKChains, and to prevent squatting/spamming
    */
    uint256 private MIN_BLOCKS_FOR_DEPOSIT_RETURN;

    /*
    * Record of total UND deposits currently in the contract
    */
    uint256 totalDeposits;

    /*
    * Mapping of the current deposits to addresses. Used for refunding
    */
    mapping(address => mapping(uint256 => uint256)) private ownerDepositMapping;

    /*
    * Mapping of ChainID -> Wrkchain Struct
    */
    mapping(uint256 => Wrkchain) public wrkchainList;

    /**
    * ---------
    * MODIFIERS
    * ---------
    */

    /*
    * @dev Modifier to ensure only current authorised WRKChain addresses can execute a function
    * @param _chainId the WRKChain network ID
    */
    modifier onlyAuth(uint256 _chainId) {
        require(wrkchainList[_chainId].isWrkchain, "ChainID does not exist");
        require(wrkchainList[_chainId].authAddresses[msg.sender] == true || msg.sender == wrkchainList[_chainId].owner);
        _;
    }

    /*
    * @dev Modifier to ensure only the WRKChain owner can execute a function
    * @param _chainId the WRKChain network ID
    */
    modifier onlyWrkchainOwner(uint256 _chainId) {
        require(wrkchainList[_chainId].isWrkchain, "ChainID does not exist");
        require(msg.sender == wrkchainList[_chainId].owner, "You are not the owner");
        _;
    }

    /**
    * -------------
    * MAIN CONTRACT
    * -------------
    **/


    /*
    * @dev contract constructor
    * @param _wrkchainRegDeposit - required deposit amount in Wei
    * @param _minBlocksForDepositReturn - int value, minimum number of blocks required to be submitted
    *                                     before deposit refund is given
    */
    function WRKChainRoot(uint256 _wrkchainRegDeposit, uint256 _minBlocksForDepositReturn) public {
        require(_wrkchainRegDeposit > 0, "Deposit amount should be > 0");
        require(_minBlocksForDepositReturn > 0, "Min blocks should be > 0");

        WRKCHAIN_REG_DEPOSIT = _wrkchainRegDeposit;
        MIN_BLOCKS_FOR_DEPOSIT_RETURN = _minBlocksForDepositReturn;
    }

    /*
    * @dev fallback function for the contract. Log event so UND can be tracked and returned
    */
    function() payable external {
        //Log who sent, and how much so it can be returned
        emit LogFallbackFunctionCalled(msg.sender, msg.value);
    }

    /*
    * @dev Register a new WRKChain.
    *      Payable - requires depost paying in order to register a new WRKChain.
    *      Deposit is refunded after defined period.
    * @param _chainId - the WRKChain Network ID
    * @param _authAddresses - list of addresses authorised to submit hashes via recordHeader
    * @param _genesisHash - WRKChain's hashed Genesis block
    */
    function registerWrkChain(
        uint256 _chainId,
        address[] _authAddresses,
        bytes32 _genesisHash) payable external {

        require(!wrkchainList[_chainId].isWrkchain, "Chain ID already exists");
        require(msg.value == WRKCHAIN_REG_DEPOSIT && msg.value > 0, "Deposit required");
        require(_chainId > 0, "Chain ID required");
        require(_genesisHash.length > 0, "Genesis hash required");
        require(_authAddresses.length > 0, "Initial Authorised addresses required");
        require(_authAddresses.length <= 10, "Maximum 10 Initial Authorised addresses allowed");


        wrkchainList[_chainId] = Wrkchain({
            lastBlock: 0,
            genesisHash: _genesisHash,
            isWrkchain: true,
            owner: msg.sender,
            numBlocksSubmitted: 0
        });

        emit RegisterWrkChain(_chainId, _genesisHash);

        for (uint i=0; i< _authAddresses.length; i++) {
            wrkchainList[_chainId].authAddresses[_authAddresses[i]] = true;
            emit AuthoriseNewAddress(_chainId, msg.sender, _authAddresses[i]);
        }

        wrkchainList[_chainId].authAddresses[msg.sender] = true;
        emit AuthoriseNewAddress(_chainId, msg.sender, msg.sender);

        totalDeposits = totalDeposits.add(msg.value);
        ownerDepositMapping[msg.sender][_chainId] = msg.value;

        emit WRKChainDepositPaid(_chainId, msg.sender, msg.value);

    }

    /*
    * @dev Record a WRKChain's block hash(es). Also handles refunding the deposit
    * A block can only be submitted once. Block data can't be overwritten
    * @param _chainId - the WRKChain Network ID
    * @param _height - Block number being submitted
    * @param _hash - WRKChain's main header hash
    * @param _parentHash - Header hash of previous block
    * @param _receiptRoot - Merkle root of WRKChain's receipts
    * @param _txRoot - Merkle root of WRKChain's Txs
    * @param _stateRoot - Merkle root of WRKChain's state DB
    * @param _blockSigner - address of the block signer (e.g. coinbase, or derived from extraData for clique etc.)
    */
    function recordHeader(
        uint256 _chainId,
        uint256 _height,
        bytes32 _hash,
        bytes32 _parentHash,
        bytes32 _receiptRoot,
        bytes32 _txRoot,
        bytes32 _stateRoot,
        address _blockSigner) public onlyAuth(_chainId) {

        require(_chainId > 0, "Chain ID required");
        require(_height > 0, "Block number required");
        require(_hash.length > 0, "Hash required");

        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        Wrkchain storage my_wrkchain = wrkchainList[_chainId];

        require(my_wrkchain.lastBlock < _height, "Block header already recorded");

        wrkchainList[_chainId].lastBlock = _height;

        wrkchainList[_chainId].numBlocksSubmitted = wrkchainList[_chainId].numBlocksSubmitted.add(1);

        // Check if ready for refund
        if(wrkchainList[_chainId].numBlocksSubmitted == MIN_BLOCKS_FOR_DEPOSIT_RETURN &&
        ownerDepositMapping[wrkchainList[_chainId].owner][_chainId] == WRKCHAIN_REG_DEPOSIT) {
            ownerDepositMapping[my_wrkchain.owner][_chainId] = 0;
            totalDeposits = totalDeposits.sub(WRKCHAIN_REG_DEPOSIT);
            emit WRKChainDepositRefund(_chainId, my_wrkchain.owner, WRKCHAIN_REG_DEPOSIT);
            my_wrkchain.owner.transfer(WRKCHAIN_REG_DEPOSIT);
        }

        emit RecordHeader(
            _chainId,
            _height,
            _hash,
            _parentHash,
            _receiptRoot,
            _txRoot,
            _stateRoot,
            _blockSigner
        );
    }

    /**
    * ---------------
    * ADMIN FUNCTIONS
    * ---------------
    */

    /*
    * @dev Authorise a new wallet to write hashes for the WRKChain.
    *      Only WRKChain owner can all this
    * @param _chainID - WRKChain network ID
    * @param _newAuthAddresses - array of one or more new addresses that can write hashes
    */
    function addAuthAddresses(
        uint256 _chainId,
        address[] memory _newAuthAddresses
    ) public onlyWrkchainOwner(_chainId) {
        require(_chainId > 0, "Chain ID required");
        require(_newAuthAddresses.length > 0, "Auth addresses required");
        require(_newAuthAddresses.length <= 10, "Max submission 10 addresses allowed");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        for (uint j=0; j< _newAuthAddresses.length; j++) {
            require(_newAuthAddresses[j] != address(0), "cannot authorise zero address");
            wrkchainList[_chainId].authAddresses[_newAuthAddresses[j]] = true;
            emit AuthoriseNewAddress(_chainId, msg.sender, _newAuthAddresses[j]);
        }
    }

    /*
    * @dev Remove a wallet from the list of authorised addresses that can write hashes for the WRKChain.
    *      Only WRKChain owner can all this
    * @param _chainID - WRKChain network ID
    * @param _addressesToRemove - array of one or more new addresses to remove
    */
    function removeAuthAddresses(
        uint256 _chainId,
        address[] memory _addressesToRemove
    ) public onlyWrkchainOwner(_chainId) {
        require(_chainId > 0, "Chain ID required");
        require(_addressesToRemove.length > 0, "Auth addresses required");
        require(_addressesToRemove.length <= 10, "Max submission 10 addresses allowed");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        for (uint j=0; j< _addressesToRemove.length; j++) {
            require(_addressesToRemove[j] != address(0), "cannot authorise zero address");
            delete wrkchainList[_chainId].authAddresses[_addressesToRemove[j]];
            emit RemoveAuthorisedAddress(_chainId, msg.sender, _addressesToRemove[j]);
        }
    }

    /*
    * @dev Change WRKChain ownership to a new wallet address.
    *      NOTE: This can only be called after the original owner has received their deposit
    *      Only WRKChain owner can all this
    * @param _chainID - WRKChain network ID
    * @param _newOwner - wallet address of new owner
    */
    function changeWrkchainOwner(
        uint256 _chainId,
        address _newOwner
    ) public onlyWrkchainOwner(_chainId) {
        require(_chainId > 0, "Chain ID required");
        require(_newOwner != address(0), "New owner required, and should be non-zero address");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        Wrkchain storage my_wrkchain = wrkchainList[_chainId];

        require(wrkchainList[_chainId].numBlocksSubmitted > MIN_BLOCKS_FOR_DEPOSIT_RETURN, "Owner cannot be changed before deposit has been returned");

        wrkchainList[_chainId].owner = _newOwner;
        wrkchainList[_chainId].authAddresses[_newOwner] = true;

        emit WRKChainOwnerChanged(_chainId, msg.sender, _newOwner);
    }

    /**
    * -------
    * GETTERS
    * -------
    */

    //get the Genesis hash for a WRKChain
    function getGenesis(uint256 _chainId) public view returns (bytes32 genesisHash_) {
        require(_chainId > 0, "Chain ID required");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        genesisHash_ = wrkchainList[_chainId].genesisHash;
    }

    // get last block number recorded for a WRKChain
    function getLastRecordedBlockNum(uint256 _chainId) public view returns (uint256 lastBlock_) {
        require(_chainId > 0, "Chain ID required");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        lastBlock_ = wrkchainList[_chainId].lastBlock;
    }

    //get number of blocks submitted
    function getNumBlocksSubmitted(uint256 _chainId) public view returns (uint256 numBlocksSubmitted_) {
        require(_chainId > 0, "Chain ID required");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        numBlocksSubmitted_ = wrkchainList[_chainId].numBlocksSubmitted;
    }

    /**
    * --------------
    * HELPER QUERIES
    * --------------
    */

    function isAuthorisedAddress(uint256 _chainId, address _address) public view returns (bool _isAuthorised) {
        require(_chainId > 0, "Chain ID required");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        _isAuthorised = wrkchainList[_chainId].authAddresses[_address];
    }

}
