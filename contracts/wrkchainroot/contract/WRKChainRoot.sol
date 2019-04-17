pragma solidity >=0.4.22 <0.6.0;

contract WRKChainRoot {

    struct Wrkchain {
        uint256 lastBlock;
        bytes32 genesisHash;
        bool isWrkchain;
        mapping(address => bool) authAddresses; //Current wallets allowed to write hashes
        address[] authAddressesIdx;
    }

    //ChainID => last block # submitted mapping
    mapping(uint256 => Wrkchain) public wrkchainList;

    mapping(address => bool) registrars;

    address masterRegistrar;

    //RecordHeader event
    event RecordHeader(
        uint256 _chainId,
        uint256 _height,
        bytes32 _hash,
        bytes32 _parentHash,
        bytes32 _receiptRoot,
        bytes32 _txRoot,
        bytes32 _stateRoot,
        address _blockSigner
    );

    event RegisterWrkChain(
        uint256 _chainId,
        bytes32 _genesisHash
    );

    //Modifier to ensure only current authorised WRKChain
    //addresses can execute a function
    modifier onlyAuth(uint256 _chainId) {
        require(wrkchainList[_chainId].isWrkchain, "ChainID does not exist");
        require(wrkchainList[_chainId].authAddresses[msg.sender] == true);
        _;
    }

    modifier onlyRegistrar() {
        require(registrars[msg.sender] == true || msg.sender == masterRegistrar);
        _;
    }

    //Contract constructor.
    constructor(address _ownerRegistrar) public
    {
        //masterRegistrar = _ownerRegistrar;
    }

    function addRegistrar(address _registrar) public onlyRegistrar {
        registrars[_registrar] = true;
    }

    function removeRegistrar(address _registrar) public onlyRegistrar {
        registrars[_registrar] = false;
    }

    function registerWrkChain(
        uint256 _chainId,
        address[] memory _authAddresses,
        bytes32 _genesisHash) public onlyRegistrar() {

        require(_chainId > 0, "Chain ID required");
        require(_genesisHash.length > 0, "Genesis hash required");
        require(_authAddresses.length > 0, "Initial Authorised addresses required");
        require(_authAddresses.length <= 10, "Maximum 10 Initial Authorised addresses allowed");

        require(!wrkchainList[_chainId].isWrkchain, "Chain ID already exists");

        wrkchainList[_chainId] = Wrkchain({
            lastBlock: 0,
            genesisHash: _genesisHash,
            isWrkchain: true,
            authAddressesIdx: _authAddresses
        });

        for (uint i=0; i< _authAddresses.length; i++) {
            wrkchainList[_chainId].authAddresses[_authAddresses[i]] = true;
        }

        emit RegisterWrkChain(_chainId, _genesisHash);

    }

    //Record a wrkchain block header
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
        require(_hash.length > 0, "Hash required");

        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        Wrkchain storage my_wrkchain = wrkchainList[_chainId];

        require(my_wrkchain.lastBlock < _height, "Block header already recorded");

        wrkchainList[_chainId].lastBlock = _height;

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

    //Set the new EVs
    function setAuthAddresses(uint256 _chainId, address[] memory _newAuthAddresses) public onlyAuth(_chainId) {
        require(_chainId > 0, "Chain ID required");
        require(_newAuthAddresses.length > 0, "Auth addresses required");
        require(_newAuthAddresses.length <= 0, "Max submission 10 addresses allowed");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        Wrkchain storage my_wrkchain = wrkchainList[_chainId];

        for(uint i =0; i < my_wrkchain.authAddressesIdx.length; i++) {
            wrkchainList[_chainId].authAddresses[my_wrkchain.authAddressesIdx[i]] = false;
        }

        for (uint i=0; i< _newAuthAddresses.length; i++) {
            wrkchainList[_chainId].authAddresses[_newAuthAddresses[i]] = true;
        }

        wrkchainList[_chainId].authAddressesIdx = _newAuthAddresses;
    }

    //get the Genesis hash for a WRKChain
    function getGenesis(uint256 _chainId) public view returns (bytes32 genesisHash_) {
        require(_chainId > 0, "Chain ID required");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        genesisHash_ = wrkchainList[_chainId].genesisHash;
    }


    //return list of current Authorised addresses for WRKChain
    function getAuthAddresses(uint256 _chainId) public view returns (address[] memory authAddressesIdx_) {
        require(_chainId > 0, "Chain ID required");
        require(wrkchainList[_chainId].isWrkchain, "Chain ID does not exist");

        authAddressesIdx_ = wrkchainList[_chainId].authAddressesIdx;
    }

}
