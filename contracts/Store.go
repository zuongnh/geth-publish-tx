// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DnaBaseDna is an auto generated low-level Go binding around an user-defined struct.
type DnaBaseDna struct {
	Id       *big.Int
	Dna      *big.Int
	Rarity   *big.Int
	IsOpened bool
}

// ZombieAttackFightScriptStruct is an auto generated low-level Go binding around an user-defined struct.
type ZombieAttackFightScriptStruct struct {
	AttackerId           *big.Int
	IsCrit               *big.Int
	Damage               *big.Int
	EnemyZombieCurrentHP *big.Int
}

// ZombieBaseZombie is an auto generated low-level Go binding around an user-defined struct.
type ZombieBaseZombie struct {
	Id             *big.Int
	Name           string
	Dna            *big.Int
	Level          uint16
	WinCount       uint16
	LossCount      uint16
	BreedCount     uint16
	Sex            uint8
	HealthPoint    uint16
	Attack         uint16
	Defense        uint16
	CriticalRate   uint16
	CriticalDamage uint16
	Speed          uint16
	CombatPower    uint16
	Rarity         string
	Exp            *big.Int
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_zombieId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_targetId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"attackerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"isCrit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"damage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"enemyZombieCurrentHP\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structZombieAttack.FightScriptStruct[]\",\"name\":\"fightScripts\",\"type\":\"tuple[]\"}],\"name\":\"FightScript\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_zombieId\",\"type\":\"uint256\"}],\"name\":\"FindBattle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dnaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rarity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NewDna\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zombieId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumZombieBase.Sex\",\"name\":\"sex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"}],\"name\":\"NewZombie\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"winCount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"lossCount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"breedCount\",\"type\":\"uint16\"},{\"internalType\":\"enumZombieBase.Sex\",\"name\":\"sex\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"healthPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"attack\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defense\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"criticalRate\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"criticalDamage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"speed\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"combatPower\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"rarity\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structZombieBase.Zombie[]\",\"name\":\"zombies\",\"type\":\"tuple[]\"}],\"name\":\"OpenStarterPack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_zombieId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerExp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loserExp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rarity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOpened\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structDnaBase.Dna\",\"name\":\"dnaSample\",\"type\":\"tuple\"}],\"name\":\"RewardUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"UserCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldLevel\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newLevel\",\"type\":\"uint16\"}],\"name\":\"UserLevelUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_zombieId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"ZombieNameChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_zombieId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetId\",\"type\":\"uint256\"}],\"name\":\"attack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fatherId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_motherId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"breedZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_zombieId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"changeZombieName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"checkOpenStarterPack\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dnaCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dnas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rarity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOpened\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dnasKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_zombieId\",\"type\":\"uint256\"}],\"name\":\"findBattle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getDnaOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rarity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOpened\",\"type\":\"bool\"}],\"internalType\":\"structDnaBase.Dna[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getZombieOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"winCount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"lossCount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"breedCount\",\"type\":\"uint16\"},{\"internalType\":\"enumZombieBase.Sex\",\"name\":\"sex\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"healthPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"attack\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defense\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"criticalRate\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"criticalDamage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"speed\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"combatPower\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"rarity\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"internalType\":\"structZombieBase.Zombie[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrateData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newFtContract\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oldFtContract\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oldNftContract\",\"outputs\":[{\"internalType\":\"contractGiftPack\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dnaId\",\"type\":\"uint256\"}],\"name\":\"openDna\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"winCount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"lossCount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"breedCount\",\"type\":\"uint16\"},{\"internalType\":\"enumZombieBase.Sex\",\"name\":\"sex\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"healthPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"attack\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defense\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"criticalRate\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"criticalDamage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"speed\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"combatPower\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"rarity\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"internalType\":\"structZombieBase.Zombie\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openStarterPack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldFt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newFt\",\"type\":\"address\"}],\"name\":\"setFtContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldNft\",\"type\":\"address\"}],\"name\":\"setNftContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"zombies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"winCount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"lossCount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"breedCount\",\"type\":\"uint16\"},{\"internalType\":\"enumZombieBase.Sex\",\"name\":\"sex\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"healthPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"attack\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defense\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"criticalRate\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"criticalDamage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"speed\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"combatPower\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"rarity\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"zombiesKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Store *StoreCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Store *StoreSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Store *StoreCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, owner)
}

// CheckOpenStarterPack is a free data retrieval call binding the contract method 0xe8c94539.
//
// Solidity: function checkOpenStarterPack(address _address) view returns(bool)
func (_Store *StoreCaller) CheckOpenStarterPack(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "checkOpenStarterPack", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckOpenStarterPack is a free data retrieval call binding the contract method 0xe8c94539.
//
// Solidity: function checkOpenStarterPack(address _address) view returns(bool)
func (_Store *StoreSession) CheckOpenStarterPack(_address common.Address) (bool, error) {
	return _Store.Contract.CheckOpenStarterPack(&_Store.CallOpts, _address)
}

// CheckOpenStarterPack is a free data retrieval call binding the contract method 0xe8c94539.
//
// Solidity: function checkOpenStarterPack(address _address) view returns(bool)
func (_Store *StoreCallerSession) CheckOpenStarterPack(_address common.Address) (bool, error) {
	return _Store.Contract.CheckOpenStarterPack(&_Store.CallOpts, _address)
}

// DnaCount is a free data retrieval call binding the contract method 0x2b596c5f.
//
// Solidity: function dnaCount() view returns(uint256)
func (_Store *StoreCaller) DnaCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "dnaCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DnaCount is a free data retrieval call binding the contract method 0x2b596c5f.
//
// Solidity: function dnaCount() view returns(uint256)
func (_Store *StoreSession) DnaCount() (*big.Int, error) {
	return _Store.Contract.DnaCount(&_Store.CallOpts)
}

// DnaCount is a free data retrieval call binding the contract method 0x2b596c5f.
//
// Solidity: function dnaCount() view returns(uint256)
func (_Store *StoreCallerSession) DnaCount() (*big.Int, error) {
	return _Store.Contract.DnaCount(&_Store.CallOpts)
}

// Dnas is a free data retrieval call binding the contract method 0x34e80ea3.
//
// Solidity: function dnas(uint256 ) view returns(uint256 id, uint256 dna, uint256 rarity, bool isOpened)
func (_Store *StoreCaller) Dnas(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id       *big.Int
	Dna      *big.Int
	Rarity   *big.Int
	IsOpened bool
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "dnas", arg0)

	outstruct := new(struct {
		Id       *big.Int
		Dna      *big.Int
		Rarity   *big.Int
		IsOpened bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Dna = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Rarity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsOpened = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Dnas is a free data retrieval call binding the contract method 0x34e80ea3.
//
// Solidity: function dnas(uint256 ) view returns(uint256 id, uint256 dna, uint256 rarity, bool isOpened)
func (_Store *StoreSession) Dnas(arg0 *big.Int) (struct {
	Id       *big.Int
	Dna      *big.Int
	Rarity   *big.Int
	IsOpened bool
}, error) {
	return _Store.Contract.Dnas(&_Store.CallOpts, arg0)
}

// Dnas is a free data retrieval call binding the contract method 0x34e80ea3.
//
// Solidity: function dnas(uint256 ) view returns(uint256 id, uint256 dna, uint256 rarity, bool isOpened)
func (_Store *StoreCallerSession) Dnas(arg0 *big.Int) (struct {
	Id       *big.Int
	Dna      *big.Int
	Rarity   *big.Int
	IsOpened bool
}, error) {
	return _Store.Contract.Dnas(&_Store.CallOpts, arg0)
}

// DnasKeys is a free data retrieval call binding the contract method 0x7e81e6de.
//
// Solidity: function dnasKeys(uint256 ) view returns(uint256)
func (_Store *StoreCaller) DnasKeys(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "dnasKeys", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DnasKeys is a free data retrieval call binding the contract method 0x7e81e6de.
//
// Solidity: function dnasKeys(uint256 ) view returns(uint256)
func (_Store *StoreSession) DnasKeys(arg0 *big.Int) (*big.Int, error) {
	return _Store.Contract.DnasKeys(&_Store.CallOpts, arg0)
}

// DnasKeys is a free data retrieval call binding the contract method 0x7e81e6de.
//
// Solidity: function dnasKeys(uint256 ) view returns(uint256)
func (_Store *StoreCallerSession) DnasKeys(arg0 *big.Int) (*big.Int, error) {
	return _Store.Contract.DnasKeys(&_Store.CallOpts, arg0)
}

// FindBattle is a free data retrieval call binding the contract method 0x68eae19f.
//
// Solidity: function findBattle(uint256 _zombieId) view returns(uint256)
func (_Store *StoreCaller) FindBattle(opts *bind.CallOpts, _zombieId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "findBattle", _zombieId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FindBattle is a free data retrieval call binding the contract method 0x68eae19f.
//
// Solidity: function findBattle(uint256 _zombieId) view returns(uint256)
func (_Store *StoreSession) FindBattle(_zombieId *big.Int) (*big.Int, error) {
	return _Store.Contract.FindBattle(&_Store.CallOpts, _zombieId)
}

// FindBattle is a free data retrieval call binding the contract method 0x68eae19f.
//
// Solidity: function findBattle(uint256 _zombieId) view returns(uint256)
func (_Store *StoreCallerSession) FindBattle(_zombieId *big.Int) (*big.Int, error) {
	return _Store.Contract.FindBattle(&_Store.CallOpts, _zombieId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Store *StoreCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Store *StoreSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.GetApproved(&_Store.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Store *StoreCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.GetApproved(&_Store.CallOpts, tokenId)
}

// GetDnaOf is a free data retrieval call binding the contract method 0xa496a79c.
//
// Solidity: function getDnaOf(address _owner) view returns((uint256,uint256,uint256,bool)[])
func (_Store *StoreCaller) GetDnaOf(opts *bind.CallOpts, _owner common.Address) ([]DnaBaseDna, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getDnaOf", _owner)

	if err != nil {
		return *new([]DnaBaseDna), err
	}

	out0 := *abi.ConvertType(out[0], new([]DnaBaseDna)).(*[]DnaBaseDna)

	return out0, err

}

// GetDnaOf is a free data retrieval call binding the contract method 0xa496a79c.
//
// Solidity: function getDnaOf(address _owner) view returns((uint256,uint256,uint256,bool)[])
func (_Store *StoreSession) GetDnaOf(_owner common.Address) ([]DnaBaseDna, error) {
	return _Store.Contract.GetDnaOf(&_Store.CallOpts, _owner)
}

// GetDnaOf is a free data retrieval call binding the contract method 0xa496a79c.
//
// Solidity: function getDnaOf(address _owner) view returns((uint256,uint256,uint256,bool)[])
func (_Store *StoreCallerSession) GetDnaOf(_owner common.Address) ([]DnaBaseDna, error) {
	return _Store.Contract.GetDnaOf(&_Store.CallOpts, _owner)
}

// GetZombieOf is a free data retrieval call binding the contract method 0x77fb99d1.
//
// Solidity: function getZombieOf(address _owner) view returns((uint256,string,uint256,uint16,uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint16,uint16,uint16,string,uint256)[])
func (_Store *StoreCaller) GetZombieOf(opts *bind.CallOpts, _owner common.Address) ([]ZombieBaseZombie, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getZombieOf", _owner)

	if err != nil {
		return *new([]ZombieBaseZombie), err
	}

	out0 := *abi.ConvertType(out[0], new([]ZombieBaseZombie)).(*[]ZombieBaseZombie)

	return out0, err

}

// GetZombieOf is a free data retrieval call binding the contract method 0x77fb99d1.
//
// Solidity: function getZombieOf(address _owner) view returns((uint256,string,uint256,uint16,uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint16,uint16,uint16,string,uint256)[])
func (_Store *StoreSession) GetZombieOf(_owner common.Address) ([]ZombieBaseZombie, error) {
	return _Store.Contract.GetZombieOf(&_Store.CallOpts, _owner)
}

// GetZombieOf is a free data retrieval call binding the contract method 0x77fb99d1.
//
// Solidity: function getZombieOf(address _owner) view returns((uint256,string,uint256,uint16,uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint16,uint16,uint16,string,uint256)[])
func (_Store *StoreCallerSession) GetZombieOf(_owner common.Address) ([]ZombieBaseZombie, error) {
	return _Store.Contract.GetZombieOf(&_Store.CallOpts, _owner)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Store *StoreCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Store *StoreSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Store.Contract.IsApprovedForAll(&_Store.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Store *StoreCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Store.Contract.IsApprovedForAll(&_Store.CallOpts, owner, operator)
}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_Store *StoreCaller) IsMigrated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "isMigrated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_Store *StoreSession) IsMigrated() (bool, error) {
	return _Store.Contract.IsMigrated(&_Store.CallOpts)
}

// IsMigrated is a free data retrieval call binding the contract method 0xb06faf62.
//
// Solidity: function isMigrated() view returns(bool)
func (_Store *StoreCallerSession) IsMigrated() (bool, error) {
	return _Store.Contract.IsMigrated(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreCallerSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// NewFtContract is a free data retrieval call binding the contract method 0x281a6478.
//
// Solidity: function newFtContract() view returns(address)
func (_Store *StoreCaller) NewFtContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "newFtContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewFtContract is a free data retrieval call binding the contract method 0x281a6478.
//
// Solidity: function newFtContract() view returns(address)
func (_Store *StoreSession) NewFtContract() (common.Address, error) {
	return _Store.Contract.NewFtContract(&_Store.CallOpts)
}

// NewFtContract is a free data retrieval call binding the contract method 0x281a6478.
//
// Solidity: function newFtContract() view returns(address)
func (_Store *StoreCallerSession) NewFtContract() (common.Address, error) {
	return _Store.Contract.NewFtContract(&_Store.CallOpts)
}

// OldFtContract is a free data retrieval call binding the contract method 0xb68d0e02.
//
// Solidity: function oldFtContract() view returns(address)
func (_Store *StoreCaller) OldFtContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "oldFtContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OldFtContract is a free data retrieval call binding the contract method 0xb68d0e02.
//
// Solidity: function oldFtContract() view returns(address)
func (_Store *StoreSession) OldFtContract() (common.Address, error) {
	return _Store.Contract.OldFtContract(&_Store.CallOpts)
}

// OldFtContract is a free data retrieval call binding the contract method 0xb68d0e02.
//
// Solidity: function oldFtContract() view returns(address)
func (_Store *StoreCallerSession) OldFtContract() (common.Address, error) {
	return _Store.Contract.OldFtContract(&_Store.CallOpts)
}

// OldNftContract is a free data retrieval call binding the contract method 0x0fb62de6.
//
// Solidity: function oldNftContract() view returns(address)
func (_Store *StoreCaller) OldNftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "oldNftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OldNftContract is a free data retrieval call binding the contract method 0x0fb62de6.
//
// Solidity: function oldNftContract() view returns(address)
func (_Store *StoreSession) OldNftContract() (common.Address, error) {
	return _Store.Contract.OldNftContract(&_Store.CallOpts)
}

// OldNftContract is a free data retrieval call binding the contract method 0x0fb62de6.
//
// Solidity: function oldNftContract() view returns(address)
func (_Store *StoreCallerSession) OldNftContract() (common.Address, error) {
	return _Store.Contract.OldNftContract(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Store *StoreCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Store *StoreSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.OwnerOf(&_Store.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Store *StoreCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.OwnerOf(&_Store.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Store.Contract.SupportsInterface(&_Store.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Store.Contract.SupportsInterface(&_Store.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreCallerSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Store *StoreCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Store *StoreSession) Token() (common.Address, error) {
	return _Store.Contract.Token(&_Store.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Store *StoreCallerSession) Token() (common.Address, error) {
	return _Store.Contract.Token(&_Store.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() view returns(uint256)
func (_Store *StoreCaller) TokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "tokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() view returns(uint256)
func (_Store *StoreSession) TokenCount() (*big.Int, error) {
	return _Store.Contract.TokenCount(&_Store.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() view returns(uint256)
func (_Store *StoreCallerSession) TokenCount() (*big.Int, error) {
	return _Store.Contract.TokenCount(&_Store.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Store *StoreCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Store *StoreSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Store.Contract.TokenURI(&_Store.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Store *StoreCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Store.Contract.TokenURI(&_Store.CallOpts, tokenId)
}

// UserCount is a free data retrieval call binding the contract method 0x07973ccf.
//
// Solidity: function userCount() view returns(uint256)
func (_Store *StoreCaller) UserCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "userCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserCount is a free data retrieval call binding the contract method 0x07973ccf.
//
// Solidity: function userCount() view returns(uint256)
func (_Store *StoreSession) UserCount() (*big.Int, error) {
	return _Store.Contract.UserCount(&_Store.CallOpts)
}

// UserCount is a free data retrieval call binding the contract method 0x07973ccf.
//
// Solidity: function userCount() view returns(uint256)
func (_Store *StoreCallerSession) UserCount() (*big.Int, error) {
	return _Store.Contract.UserCount(&_Store.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(uint256 id, address walletAddress, uint16 level, uint256 exp)
func (_Store *StoreCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id            *big.Int
	WalletAddress common.Address
	Level         uint16
	Exp           *big.Int
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		Id            *big.Int
		WalletAddress common.Address
		Level         uint16
		Exp           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WalletAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Level = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Exp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(uint256 id, address walletAddress, uint16 level, uint256 exp)
func (_Store *StoreSession) Users(arg0 *big.Int) (struct {
	Id            *big.Int
	WalletAddress common.Address
	Level         uint16
	Exp           *big.Int
}, error) {
	return _Store.Contract.Users(&_Store.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(uint256 id, address walletAddress, uint16 level, uint256 exp)
func (_Store *StoreCallerSession) Users(arg0 *big.Int) (struct {
	Id            *big.Int
	WalletAddress common.Address
	Level         uint16
	Exp           *big.Int
}, error) {
	return _Store.Contract.Users(&_Store.CallOpts, arg0)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_Store *StoreCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_Store *StoreSession) ZombieCount() (*big.Int, error) {
	return _Store.Contract.ZombieCount(&_Store.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_Store *StoreCallerSession) ZombieCount() (*big.Int, error) {
	return _Store.Contract.ZombieCount(&_Store.CallOpts)
}

// Zombies is a free data retrieval call binding the contract method 0x2052465e.
//
// Solidity: function zombies(uint256 ) view returns(uint256 id, string name, uint256 dna, uint16 level, uint16 winCount, uint16 lossCount, uint16 breedCount, uint8 sex, uint16 healthPoint, uint16 attack, uint16 defense, uint16 criticalRate, uint16 criticalDamage, uint16 speed, uint16 combatPower, string rarity, uint256 exp)
func (_Store *StoreCaller) Zombies(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	Dna            *big.Int
	Level          uint16
	WinCount       uint16
	LossCount      uint16
	BreedCount     uint16
	Sex            uint8
	HealthPoint    uint16
	Attack         uint16
	Defense        uint16
	CriticalRate   uint16
	CriticalDamage uint16
	Speed          uint16
	CombatPower    uint16
	Rarity         string
	Exp            *big.Int
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "zombies", arg0)

	outstruct := new(struct {
		Id             *big.Int
		Name           string
		Dna            *big.Int
		Level          uint16
		WinCount       uint16
		LossCount      uint16
		BreedCount     uint16
		Sex            uint8
		HealthPoint    uint16
		Attack         uint16
		Defense        uint16
		CriticalRate   uint16
		CriticalDamage uint16
		Speed          uint16
		CombatPower    uint16
		Rarity         string
		Exp            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Dna = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Level = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.WinCount = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.LossCount = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.BreedCount = *abi.ConvertType(out[6], new(uint16)).(*uint16)
	outstruct.Sex = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.HealthPoint = *abi.ConvertType(out[8], new(uint16)).(*uint16)
	outstruct.Attack = *abi.ConvertType(out[9], new(uint16)).(*uint16)
	outstruct.Defense = *abi.ConvertType(out[10], new(uint16)).(*uint16)
	outstruct.CriticalRate = *abi.ConvertType(out[11], new(uint16)).(*uint16)
	outstruct.CriticalDamage = *abi.ConvertType(out[12], new(uint16)).(*uint16)
	outstruct.Speed = *abi.ConvertType(out[13], new(uint16)).(*uint16)
	outstruct.CombatPower = *abi.ConvertType(out[14], new(uint16)).(*uint16)
	outstruct.Rarity = *abi.ConvertType(out[15], new(string)).(*string)
	outstruct.Exp = *abi.ConvertType(out[16], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Zombies is a free data retrieval call binding the contract method 0x2052465e.
//
// Solidity: function zombies(uint256 ) view returns(uint256 id, string name, uint256 dna, uint16 level, uint16 winCount, uint16 lossCount, uint16 breedCount, uint8 sex, uint16 healthPoint, uint16 attack, uint16 defense, uint16 criticalRate, uint16 criticalDamage, uint16 speed, uint16 combatPower, string rarity, uint256 exp)
func (_Store *StoreSession) Zombies(arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	Dna            *big.Int
	Level          uint16
	WinCount       uint16
	LossCount      uint16
	BreedCount     uint16
	Sex            uint8
	HealthPoint    uint16
	Attack         uint16
	Defense        uint16
	CriticalRate   uint16
	CriticalDamage uint16
	Speed          uint16
	CombatPower    uint16
	Rarity         string
	Exp            *big.Int
}, error) {
	return _Store.Contract.Zombies(&_Store.CallOpts, arg0)
}

// Zombies is a free data retrieval call binding the contract method 0x2052465e.
//
// Solidity: function zombies(uint256 ) view returns(uint256 id, string name, uint256 dna, uint16 level, uint16 winCount, uint16 lossCount, uint16 breedCount, uint8 sex, uint16 healthPoint, uint16 attack, uint16 defense, uint16 criticalRate, uint16 criticalDamage, uint16 speed, uint16 combatPower, string rarity, uint256 exp)
func (_Store *StoreCallerSession) Zombies(arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	Dna            *big.Int
	Level          uint16
	WinCount       uint16
	LossCount      uint16
	BreedCount     uint16
	Sex            uint8
	HealthPoint    uint16
	Attack         uint16
	Defense        uint16
	CriticalRate   uint16
	CriticalDamage uint16
	Speed          uint16
	CombatPower    uint16
	Rarity         string
	Exp            *big.Int
}, error) {
	return _Store.Contract.Zombies(&_Store.CallOpts, arg0)
}

// ZombiesKeys is a free data retrieval call binding the contract method 0xb7287d90.
//
// Solidity: function zombiesKeys(uint256 ) view returns(uint256)
func (_Store *StoreCaller) ZombiesKeys(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "zombiesKeys", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombiesKeys is a free data retrieval call binding the contract method 0xb7287d90.
//
// Solidity: function zombiesKeys(uint256 ) view returns(uint256)
func (_Store *StoreSession) ZombiesKeys(arg0 *big.Int) (*big.Int, error) {
	return _Store.Contract.ZombiesKeys(&_Store.CallOpts, arg0)
}

// ZombiesKeys is a free data retrieval call binding the contract method 0xb7287d90.
//
// Solidity: function zombiesKeys(uint256 ) view returns(uint256)
func (_Store *StoreCallerSession) ZombiesKeys(arg0 *big.Int) (*big.Int, error) {
	return _Store.Contract.ZombiesKeys(&_Store.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Store *StoreTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Store *StoreSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, to, tokenId)
}

// Attack is a paid mutator transaction binding the contract method 0xe1fa7638.
//
// Solidity: function attack(uint256 _zombieId, uint256 _targetId) returns()
func (_Store *StoreTransactor) Attack(opts *bind.TransactOpts, _zombieId *big.Int, _targetId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "attack", _zombieId, _targetId)
}

// Attack is a paid mutator transaction binding the contract method 0xe1fa7638.
//
// Solidity: function attack(uint256 _zombieId, uint256 _targetId) returns()
func (_Store *StoreSession) Attack(_zombieId *big.Int, _targetId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Attack(&_Store.TransactOpts, _zombieId, _targetId)
}

// Attack is a paid mutator transaction binding the contract method 0xe1fa7638.
//
// Solidity: function attack(uint256 _zombieId, uint256 _targetId) returns()
func (_Store *StoreTransactorSession) Attack(_zombieId *big.Int, _targetId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Attack(&_Store.TransactOpts, _zombieId, _targetId)
}

// BreedZombie is a paid mutator transaction binding the contract method 0x9fba18a7.
//
// Solidity: function breedZombie(uint256 _fatherId, uint256 _motherId, string _name) returns()
func (_Store *StoreTransactor) BreedZombie(opts *bind.TransactOpts, _fatherId *big.Int, _motherId *big.Int, _name string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "breedZombie", _fatherId, _motherId, _name)
}

// BreedZombie is a paid mutator transaction binding the contract method 0x9fba18a7.
//
// Solidity: function breedZombie(uint256 _fatherId, uint256 _motherId, string _name) returns()
func (_Store *StoreSession) BreedZombie(_fatherId *big.Int, _motherId *big.Int, _name string) (*types.Transaction, error) {
	return _Store.Contract.BreedZombie(&_Store.TransactOpts, _fatherId, _motherId, _name)
}

// BreedZombie is a paid mutator transaction binding the contract method 0x9fba18a7.
//
// Solidity: function breedZombie(uint256 _fatherId, uint256 _motherId, string _name) returns()
func (_Store *StoreTransactorSession) BreedZombie(_fatherId *big.Int, _motherId *big.Int, _name string) (*types.Transaction, error) {
	return _Store.Contract.BreedZombie(&_Store.TransactOpts, _fatherId, _motherId, _name)
}

// ChangeZombieName is a paid mutator transaction binding the contract method 0x1048a791.
//
// Solidity: function changeZombieName(uint256 _zombieId, string _name) returns()
func (_Store *StoreTransactor) ChangeZombieName(opts *bind.TransactOpts, _zombieId *big.Int, _name string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "changeZombieName", _zombieId, _name)
}

// ChangeZombieName is a paid mutator transaction binding the contract method 0x1048a791.
//
// Solidity: function changeZombieName(uint256 _zombieId, string _name) returns()
func (_Store *StoreSession) ChangeZombieName(_zombieId *big.Int, _name string) (*types.Transaction, error) {
	return _Store.Contract.ChangeZombieName(&_Store.TransactOpts, _zombieId, _name)
}

// ChangeZombieName is a paid mutator transaction binding the contract method 0x1048a791.
//
// Solidity: function changeZombieName(uint256 _zombieId, string _name) returns()
func (_Store *StoreTransactorSession) ChangeZombieName(_zombieId *big.Int, _name string) (*types.Transaction, error) {
	return _Store.Contract.ChangeZombieName(&_Store.TransactOpts, _zombieId, _name)
}

// MigrateData is a paid mutator transaction binding the contract method 0x21ca8f61.
//
// Solidity: function migrateData() returns()
func (_Store *StoreTransactor) MigrateData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "migrateData")
}

// MigrateData is a paid mutator transaction binding the contract method 0x21ca8f61.
//
// Solidity: function migrateData() returns()
func (_Store *StoreSession) MigrateData() (*types.Transaction, error) {
	return _Store.Contract.MigrateData(&_Store.TransactOpts)
}

// MigrateData is a paid mutator transaction binding the contract method 0x21ca8f61.
//
// Solidity: function migrateData() returns()
func (_Store *StoreTransactorSession) MigrateData() (*types.Transaction, error) {
	return _Store.Contract.MigrateData(&_Store.TransactOpts)
}

// OpenDna is a paid mutator transaction binding the contract method 0xdf0bf6b3.
//
// Solidity: function openDna(uint256 _dnaId) returns((uint256,string,uint256,uint16,uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint16,uint16,uint16,string,uint256))
func (_Store *StoreTransactor) OpenDna(opts *bind.TransactOpts, _dnaId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "openDna", _dnaId)
}

// OpenDna is a paid mutator transaction binding the contract method 0xdf0bf6b3.
//
// Solidity: function openDna(uint256 _dnaId) returns((uint256,string,uint256,uint16,uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint16,uint16,uint16,string,uint256))
func (_Store *StoreSession) OpenDna(_dnaId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.OpenDna(&_Store.TransactOpts, _dnaId)
}

// OpenDna is a paid mutator transaction binding the contract method 0xdf0bf6b3.
//
// Solidity: function openDna(uint256 _dnaId) returns((uint256,string,uint256,uint16,uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint16,uint16,uint16,string,uint256))
func (_Store *StoreTransactorSession) OpenDna(_dnaId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.OpenDna(&_Store.TransactOpts, _dnaId)
}

// OpenStarterPack is a paid mutator transaction binding the contract method 0x78a25d95.
//
// Solidity: function openStarterPack() returns()
func (_Store *StoreTransactor) OpenStarterPack(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "openStarterPack")
}

// OpenStarterPack is a paid mutator transaction binding the contract method 0x78a25d95.
//
// Solidity: function openStarterPack() returns()
func (_Store *StoreSession) OpenStarterPack() (*types.Transaction, error) {
	return _Store.Contract.OpenStarterPack(&_Store.TransactOpts)
}

// OpenStarterPack is a paid mutator transaction binding the contract method 0x78a25d95.
//
// Solidity: function openStarterPack() returns()
func (_Store *StoreTransactorSession) OpenStarterPack() (*types.Transaction, error) {
	return _Store.Contract.OpenStarterPack(&_Store.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Store *StoreTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Store *StoreSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom0(&_Store.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Store *StoreTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom0(&_Store.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Store *StoreTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Store *StoreSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Store.Contract.SetApprovalForAll(&_Store.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Store *StoreTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Store.Contract.SetApprovalForAll(&_Store.TransactOpts, operator, approved)
}

// SetFtContract is a paid mutator transaction binding the contract method 0xf4f9653b.
//
// Solidity: function setFtContract(address _oldFt, address _newFt) returns()
func (_Store *StoreTransactor) SetFtContract(opts *bind.TransactOpts, _oldFt common.Address, _newFt common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setFtContract", _oldFt, _newFt)
}

// SetFtContract is a paid mutator transaction binding the contract method 0xf4f9653b.
//
// Solidity: function setFtContract(address _oldFt, address _newFt) returns()
func (_Store *StoreSession) SetFtContract(_oldFt common.Address, _newFt common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetFtContract(&_Store.TransactOpts, _oldFt, _newFt)
}

// SetFtContract is a paid mutator transaction binding the contract method 0xf4f9653b.
//
// Solidity: function setFtContract(address _oldFt, address _newFt) returns()
func (_Store *StoreTransactorSession) SetFtContract(_oldFt common.Address, _newFt common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetFtContract(&_Store.TransactOpts, _oldFt, _newFt)
}

// SetNftContract is a paid mutator transaction binding the contract method 0x52f5ad77.
//
// Solidity: function setNftContract(address _oldNft) returns()
func (_Store *StoreTransactor) SetNftContract(opts *bind.TransactOpts, _oldNft common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setNftContract", _oldNft)
}

// SetNftContract is a paid mutator transaction binding the contract method 0x52f5ad77.
//
// Solidity: function setNftContract(address _oldNft) returns()
func (_Store *StoreSession) SetNftContract(_oldNft common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetNftContract(&_Store.TransactOpts, _oldNft)
}

// SetNftContract is a paid mutator transaction binding the contract method 0x52f5ad77.
//
// Solidity: function setNftContract(address _oldNft) returns()
func (_Store *StoreTransactorSession) SetNftContract(_oldNft common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetNftContract(&_Store.TransactOpts, _oldNft)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Store *StoreTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Store *StoreSession) Withdraw() (*types.Transaction, error) {
	return _Store.Contract.Withdraw(&_Store.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Store *StoreTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Store.Contract.Withdraw(&_Store.TransactOpts)
}

// StoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Store contract.
type StoreApprovalIterator struct {
	Event *StoreApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreApproval represents a Approval event raised by the Store contract.
type StoreApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Store *StoreFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*StoreApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoreApprovalIterator{contract: _Store.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Store *StoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StoreApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreApproval)
				if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Store *StoreFilterer) ParseApproval(log types.Log) (*StoreApproval, error) {
	event := new(StoreApproval)
	if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Store contract.
type StoreApprovalForAllIterator struct {
	Event *StoreApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreApprovalForAll represents a ApprovalForAll event raised by the Store contract.
type StoreApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Store *StoreFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*StoreApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StoreApprovalForAllIterator{contract: _Store.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Store *StoreFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StoreApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreApprovalForAll)
				if err := _Store.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Store *StoreFilterer) ParseApprovalForAll(log types.Log) (*StoreApprovalForAll, error) {
	event := new(StoreApprovalForAll)
	if err := _Store.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreFightScriptIterator is returned from FilterFightScript and is used to iterate over the raw logs and unpacked data for FightScript events raised by the Store contract.
type StoreFightScriptIterator struct {
	Event *StoreFightScript // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreFightScriptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreFightScript)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreFightScript)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreFightScriptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreFightScriptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreFightScript represents a FightScript event raised by the Store contract.
type StoreFightScript struct {
	ZombieId     *big.Int
	TargetId     *big.Int
	FightScripts []ZombieAttackFightScriptStruct
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFightScript is a free log retrieval operation binding the contract event 0x08e991b8a3790d7db505d1129eef2d000683c010c238ebb5fa654827ddfb4add.
//
// Solidity: event FightScript(uint256 _zombieId, uint256 _targetId, (uint256,uint256,uint256,uint256)[] fightScripts)
func (_Store *StoreFilterer) FilterFightScript(opts *bind.FilterOpts) (*StoreFightScriptIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "FightScript")
	if err != nil {
		return nil, err
	}
	return &StoreFightScriptIterator{contract: _Store.contract, event: "FightScript", logs: logs, sub: sub}, nil
}

// WatchFightScript is a free log subscription operation binding the contract event 0x08e991b8a3790d7db505d1129eef2d000683c010c238ebb5fa654827ddfb4add.
//
// Solidity: event FightScript(uint256 _zombieId, uint256 _targetId, (uint256,uint256,uint256,uint256)[] fightScripts)
func (_Store *StoreFilterer) WatchFightScript(opts *bind.WatchOpts, sink chan<- *StoreFightScript) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "FightScript")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreFightScript)
				if err := _Store.contract.UnpackLog(event, "FightScript", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFightScript is a log parse operation binding the contract event 0x08e991b8a3790d7db505d1129eef2d000683c010c238ebb5fa654827ddfb4add.
//
// Solidity: event FightScript(uint256 _zombieId, uint256 _targetId, (uint256,uint256,uint256,uint256)[] fightScripts)
func (_Store *StoreFilterer) ParseFightScript(log types.Log) (*StoreFightScript, error) {
	event := new(StoreFightScript)
	if err := _Store.contract.UnpackLog(event, "FightScript", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreFindBattleIterator is returned from FilterFindBattle and is used to iterate over the raw logs and unpacked data for FindBattle events raised by the Store contract.
type StoreFindBattleIterator struct {
	Event *StoreFindBattle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreFindBattleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreFindBattle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreFindBattle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreFindBattleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreFindBattleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreFindBattle represents a FindBattle event raised by the Store contract.
type StoreFindBattle struct {
	ZombieId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFindBattle is a free log retrieval operation binding the contract event 0x8c6ddffc5324a95f2f2423e1d9146b19be5e6e3972d9a6f198abe563b66b61d7.
//
// Solidity: event FindBattle(uint256 _zombieId)
func (_Store *StoreFilterer) FilterFindBattle(opts *bind.FilterOpts) (*StoreFindBattleIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "FindBattle")
	if err != nil {
		return nil, err
	}
	return &StoreFindBattleIterator{contract: _Store.contract, event: "FindBattle", logs: logs, sub: sub}, nil
}

// WatchFindBattle is a free log subscription operation binding the contract event 0x8c6ddffc5324a95f2f2423e1d9146b19be5e6e3972d9a6f198abe563b66b61d7.
//
// Solidity: event FindBattle(uint256 _zombieId)
func (_Store *StoreFilterer) WatchFindBattle(opts *bind.WatchOpts, sink chan<- *StoreFindBattle) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "FindBattle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreFindBattle)
				if err := _Store.contract.UnpackLog(event, "FindBattle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFindBattle is a log parse operation binding the contract event 0x8c6ddffc5324a95f2f2423e1d9146b19be5e6e3972d9a6f198abe563b66b61d7.
//
// Solidity: event FindBattle(uint256 _zombieId)
func (_Store *StoreFilterer) ParseFindBattle(log types.Log) (*StoreFindBattle, error) {
	event := new(StoreFindBattle)
	if err := _Store.contract.UnpackLog(event, "FindBattle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNewDnaIterator is returned from FilterNewDna and is used to iterate over the raw logs and unpacked data for NewDna events raised by the Store contract.
type StoreNewDnaIterator struct {
	Event *StoreNewDna // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreNewDnaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNewDna)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreNewDna)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreNewDnaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNewDnaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNewDna represents a NewDna event raised by the Store contract.
type StoreNewDna struct {
	DnaId  *big.Int
	Dna    *big.Int
	Rarity *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewDna is a free log retrieval operation binding the contract event 0xada904ed15f87a9bdae46cb36c3cdc3481d3ce2cf4c55a697cb0323f2f6172dd.
//
// Solidity: event NewDna(uint256 dnaId, uint256 dna, uint256 rarity, address sender)
func (_Store *StoreFilterer) FilterNewDna(opts *bind.FilterOpts) (*StoreNewDnaIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NewDna")
	if err != nil {
		return nil, err
	}
	return &StoreNewDnaIterator{contract: _Store.contract, event: "NewDna", logs: logs, sub: sub}, nil
}

// WatchNewDna is a free log subscription operation binding the contract event 0xada904ed15f87a9bdae46cb36c3cdc3481d3ce2cf4c55a697cb0323f2f6172dd.
//
// Solidity: event NewDna(uint256 dnaId, uint256 dna, uint256 rarity, address sender)
func (_Store *StoreFilterer) WatchNewDna(opts *bind.WatchOpts, sink chan<- *StoreNewDna) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NewDna")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNewDna)
				if err := _Store.contract.UnpackLog(event, "NewDna", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewDna is a log parse operation binding the contract event 0xada904ed15f87a9bdae46cb36c3cdc3481d3ce2cf4c55a697cb0323f2f6172dd.
//
// Solidity: event NewDna(uint256 dnaId, uint256 dna, uint256 rarity, address sender)
func (_Store *StoreFilterer) ParseNewDna(log types.Log) (*StoreNewDna, error) {
	event := new(StoreNewDna)
	if err := _Store.contract.UnpackLog(event, "NewDna", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNewZombieIterator is returned from FilterNewZombie and is used to iterate over the raw logs and unpacked data for NewZombie events raised by the Store contract.
type StoreNewZombieIterator struct {
	Event *StoreNewZombie // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreNewZombieIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNewZombie)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreNewZombie)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreNewZombieIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNewZombieIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNewZombie represents a NewZombie event raised by the Store contract.
type StoreNewZombie struct {
	Sender   common.Address
	ZombieId *big.Int
	Name     string
	Dna      *big.Int
	Sex      uint8
	Level    uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewZombie is a free log retrieval operation binding the contract event 0x216e9ffdf114e85cfb59401bc6cc3c1273c51a6ea9ee514c9e0384904b67c343.
//
// Solidity: event NewZombie(address sender, uint256 zombieId, string name, uint256 dna, uint8 sex, uint16 level)
func (_Store *StoreFilterer) FilterNewZombie(opts *bind.FilterOpts) (*StoreNewZombieIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NewZombie")
	if err != nil {
		return nil, err
	}
	return &StoreNewZombieIterator{contract: _Store.contract, event: "NewZombie", logs: logs, sub: sub}, nil
}

// WatchNewZombie is a free log subscription operation binding the contract event 0x216e9ffdf114e85cfb59401bc6cc3c1273c51a6ea9ee514c9e0384904b67c343.
//
// Solidity: event NewZombie(address sender, uint256 zombieId, string name, uint256 dna, uint8 sex, uint16 level)
func (_Store *StoreFilterer) WatchNewZombie(opts *bind.WatchOpts, sink chan<- *StoreNewZombie) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NewZombie")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNewZombie)
				if err := _Store.contract.UnpackLog(event, "NewZombie", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewZombie is a log parse operation binding the contract event 0x216e9ffdf114e85cfb59401bc6cc3c1273c51a6ea9ee514c9e0384904b67c343.
//
// Solidity: event NewZombie(address sender, uint256 zombieId, string name, uint256 dna, uint8 sex, uint16 level)
func (_Store *StoreFilterer) ParseNewZombie(log types.Log) (*StoreNewZombie, error) {
	event := new(StoreNewZombie)
	if err := _Store.contract.UnpackLog(event, "NewZombie", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOpenStarterPackIterator is returned from FilterOpenStarterPack and is used to iterate over the raw logs and unpacked data for OpenStarterPack events raised by the Store contract.
type StoreOpenStarterPackIterator struct {
	Event *StoreOpenStarterPack // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreOpenStarterPackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOpenStarterPack)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreOpenStarterPack)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreOpenStarterPackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOpenStarterPackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOpenStarterPack represents a OpenStarterPack event raised by the Store contract.
type StoreOpenStarterPack struct {
	Owner   common.Address
	Zombies []ZombieBaseZombie
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpenStarterPack is a free log retrieval operation binding the contract event 0x2f543ab1f714c87443f70c5f393d3dbdac75f95e461996ff1243aab5d8c40b26.
//
// Solidity: event OpenStarterPack(address indexed owner, (uint256,string,uint256,uint16,uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint16,uint16,uint16,string,uint256)[] zombies)
func (_Store *StoreFilterer) FilterOpenStarterPack(opts *bind.FilterOpts, owner []common.Address) (*StoreOpenStarterPackIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "OpenStarterPack", ownerRule)
	if err != nil {
		return nil, err
	}
	return &StoreOpenStarterPackIterator{contract: _Store.contract, event: "OpenStarterPack", logs: logs, sub: sub}, nil
}

// WatchOpenStarterPack is a free log subscription operation binding the contract event 0x2f543ab1f714c87443f70c5f393d3dbdac75f95e461996ff1243aab5d8c40b26.
//
// Solidity: event OpenStarterPack(address indexed owner, (uint256,string,uint256,uint16,uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint16,uint16,uint16,string,uint256)[] zombies)
func (_Store *StoreFilterer) WatchOpenStarterPack(opts *bind.WatchOpts, sink chan<- *StoreOpenStarterPack, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "OpenStarterPack", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOpenStarterPack)
				if err := _Store.contract.UnpackLog(event, "OpenStarterPack", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOpenStarterPack is a log parse operation binding the contract event 0x2f543ab1f714c87443f70c5f393d3dbdac75f95e461996ff1243aab5d8c40b26.
//
// Solidity: event OpenStarterPack(address indexed owner, (uint256,string,uint256,uint16,uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint16,uint16,uint16,string,uint256)[] zombies)
func (_Store *StoreFilterer) ParseOpenStarterPack(log types.Log) (*StoreOpenStarterPack, error) {
	event := new(StoreOpenStarterPack)
	if err := _Store.contract.UnpackLog(event, "OpenStarterPack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Store contract.
type StoreOwnershipTransferredIterator struct {
	Event *StoreOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOwnershipTransferred represents a OwnershipTransferred event raised by the Store contract.
type StoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Store *StoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StoreOwnershipTransferredIterator{contract: _Store.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Store *StoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOwnershipTransferred)
				if err := _Store.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Store *StoreFilterer) ParseOwnershipTransferred(log types.Log) (*StoreOwnershipTransferred, error) {
	event := new(StoreOwnershipTransferred)
	if err := _Store.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRewardUserIterator is returned from FilterRewardUser and is used to iterate over the raw logs and unpacked data for RewardUser events raised by the Store contract.
type StoreRewardUserIterator struct {
	Event *StoreRewardUser // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreRewardUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRewardUser)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreRewardUser)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreRewardUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRewardUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRewardUser represents a RewardUser event raised by the Store contract.
type StoreRewardUser struct {
	ZombieId  *big.Int
	Amount    *big.Int
	WinnerExp *big.Int
	LoserExp  *big.Int
	DnaSample DnaBaseDna
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardUser is a free log retrieval operation binding the contract event 0xbd1baa7eb371fe183d9896e676baff51c8b7bf46a5cb7928c787bf916a1b888b.
//
// Solidity: event RewardUser(uint256 _zombieId, uint256 amount, uint256 winnerExp, uint256 loserExp, (uint256,uint256,uint256,bool) dnaSample)
func (_Store *StoreFilterer) FilterRewardUser(opts *bind.FilterOpts) (*StoreRewardUserIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "RewardUser")
	if err != nil {
		return nil, err
	}
	return &StoreRewardUserIterator{contract: _Store.contract, event: "RewardUser", logs: logs, sub: sub}, nil
}

// WatchRewardUser is a free log subscription operation binding the contract event 0xbd1baa7eb371fe183d9896e676baff51c8b7bf46a5cb7928c787bf916a1b888b.
//
// Solidity: event RewardUser(uint256 _zombieId, uint256 amount, uint256 winnerExp, uint256 loserExp, (uint256,uint256,uint256,bool) dnaSample)
func (_Store *StoreFilterer) WatchRewardUser(opts *bind.WatchOpts, sink chan<- *StoreRewardUser) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "RewardUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRewardUser)
				if err := _Store.contract.UnpackLog(event, "RewardUser", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardUser is a log parse operation binding the contract event 0xbd1baa7eb371fe183d9896e676baff51c8b7bf46a5cb7928c787bf916a1b888b.
//
// Solidity: event RewardUser(uint256 _zombieId, uint256 amount, uint256 winnerExp, uint256 loserExp, (uint256,uint256,uint256,bool) dnaSample)
func (_Store *StoreFilterer) ParseRewardUser(log types.Log) (*StoreRewardUser, error) {
	event := new(StoreRewardUser)
	if err := _Store.contract.UnpackLog(event, "RewardUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Store contract.
type StoreTransferIterator struct {
	Event *StoreTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransfer represents a Transfer event raised by the Store contract.
type StoreTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Store *StoreFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*StoreTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoreTransferIterator{contract: _Store.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Store *StoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoreTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransfer)
				if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Store *StoreFilterer) ParseTransfer(log types.Log) (*StoreTransfer, error) {
	event := new(StoreTransfer)
	if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreUserCreatedIterator is returned from FilterUserCreated and is used to iterate over the raw logs and unpacked data for UserCreated events raised by the Store contract.
type StoreUserCreatedIterator struct {
	Event *StoreUserCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreUserCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreUserCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreUserCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreUserCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreUserCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreUserCreated represents a UserCreated event raised by the Store contract.
type StoreUserCreated struct {
	Sender common.Address
	Level  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUserCreated is a free log retrieval operation binding the contract event 0x855e943c5f44cb7dcc68d332cd0bfa8275cc4ea725648d4c53a6bc3ae7df984b.
//
// Solidity: event UserCreated(address sender, uint8 level)
func (_Store *StoreFilterer) FilterUserCreated(opts *bind.FilterOpts) (*StoreUserCreatedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "UserCreated")
	if err != nil {
		return nil, err
	}
	return &StoreUserCreatedIterator{contract: _Store.contract, event: "UserCreated", logs: logs, sub: sub}, nil
}

// WatchUserCreated is a free log subscription operation binding the contract event 0x855e943c5f44cb7dcc68d332cd0bfa8275cc4ea725648d4c53a6bc3ae7df984b.
//
// Solidity: event UserCreated(address sender, uint8 level)
func (_Store *StoreFilterer) WatchUserCreated(opts *bind.WatchOpts, sink chan<- *StoreUserCreated) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "UserCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreUserCreated)
				if err := _Store.contract.UnpackLog(event, "UserCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserCreated is a log parse operation binding the contract event 0x855e943c5f44cb7dcc68d332cd0bfa8275cc4ea725648d4c53a6bc3ae7df984b.
//
// Solidity: event UserCreated(address sender, uint8 level)
func (_Store *StoreFilterer) ParseUserCreated(log types.Log) (*StoreUserCreated, error) {
	event := new(StoreUserCreated)
	if err := _Store.contract.UnpackLog(event, "UserCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreUserLevelUpIterator is returned from FilterUserLevelUp and is used to iterate over the raw logs and unpacked data for UserLevelUp events raised by the Store contract.
type StoreUserLevelUpIterator struct {
	Event *StoreUserLevelUp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreUserLevelUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreUserLevelUp)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreUserLevelUp)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreUserLevelUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreUserLevelUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreUserLevelUp represents a UserLevelUp event raised by the Store contract.
type StoreUserLevelUp struct {
	Sender   common.Address
	OldLevel uint16
	NewLevel uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUserLevelUp is a free log retrieval operation binding the contract event 0x8dd8f3aa2403fcd08819e3a9b4540580231f0a08516aa841ce1d213e4f1ae88d.
//
// Solidity: event UserLevelUp(address sender, uint16 oldLevel, uint16 newLevel)
func (_Store *StoreFilterer) FilterUserLevelUp(opts *bind.FilterOpts) (*StoreUserLevelUpIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "UserLevelUp")
	if err != nil {
		return nil, err
	}
	return &StoreUserLevelUpIterator{contract: _Store.contract, event: "UserLevelUp", logs: logs, sub: sub}, nil
}

// WatchUserLevelUp is a free log subscription operation binding the contract event 0x8dd8f3aa2403fcd08819e3a9b4540580231f0a08516aa841ce1d213e4f1ae88d.
//
// Solidity: event UserLevelUp(address sender, uint16 oldLevel, uint16 newLevel)
func (_Store *StoreFilterer) WatchUserLevelUp(opts *bind.WatchOpts, sink chan<- *StoreUserLevelUp) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "UserLevelUp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreUserLevelUp)
				if err := _Store.contract.UnpackLog(event, "UserLevelUp", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserLevelUp is a log parse operation binding the contract event 0x8dd8f3aa2403fcd08819e3a9b4540580231f0a08516aa841ce1d213e4f1ae88d.
//
// Solidity: event UserLevelUp(address sender, uint16 oldLevel, uint16 newLevel)
func (_Store *StoreFilterer) ParseUserLevelUp(log types.Log) (*StoreUserLevelUp, error) {
	event := new(StoreUserLevelUp)
	if err := _Store.contract.UnpackLog(event, "UserLevelUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreZombieNameChangedIterator is returned from FilterZombieNameChanged and is used to iterate over the raw logs and unpacked data for ZombieNameChanged events raised by the Store contract.
type StoreZombieNameChangedIterator struct {
	Event *StoreZombieNameChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreZombieNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreZombieNameChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreZombieNameChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreZombieNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreZombieNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreZombieNameChanged represents a ZombieNameChanged event raised by the Store contract.
type StoreZombieNameChanged struct {
	ZombieId *big.Int
	NewName  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterZombieNameChanged is a free log retrieval operation binding the contract event 0x299c3dddc764feba577527e74acfdf939f99aede7c8621b84616a677681f0cc3.
//
// Solidity: event ZombieNameChanged(uint256 _zombieId, string newName)
func (_Store *StoreFilterer) FilterZombieNameChanged(opts *bind.FilterOpts) (*StoreZombieNameChangedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ZombieNameChanged")
	if err != nil {
		return nil, err
	}
	return &StoreZombieNameChangedIterator{contract: _Store.contract, event: "ZombieNameChanged", logs: logs, sub: sub}, nil
}

// WatchZombieNameChanged is a free log subscription operation binding the contract event 0x299c3dddc764feba577527e74acfdf939f99aede7c8621b84616a677681f0cc3.
//
// Solidity: event ZombieNameChanged(uint256 _zombieId, string newName)
func (_Store *StoreFilterer) WatchZombieNameChanged(opts *bind.WatchOpts, sink chan<- *StoreZombieNameChanged) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ZombieNameChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreZombieNameChanged)
				if err := _Store.contract.UnpackLog(event, "ZombieNameChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZombieNameChanged is a log parse operation binding the contract event 0x299c3dddc764feba577527e74acfdf939f99aede7c8621b84616a677681f0cc3.
//
// Solidity: event ZombieNameChanged(uint256 _zombieId, string newName)
func (_Store *StoreFilterer) ParseZombieNameChanged(log types.Log) (*StoreZombieNameChanged, error) {
	event := new(StoreZombieNameChanged)
	if err := _Store.contract.UnpackLog(event, "ZombieNameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
