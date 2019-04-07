package dots

import (
	"crypto/rand"
	"github.com/fraudmarc/go-monero/walletrpc"
	"log"
	"math/big"
)

type Donation struct {
	Name    string
	Address string
}

func ChooseDonation() Donation {
	var donations []Donation

	if config.Mainnet == false {
		donations = append(donations,
			Donation{
				Name:    "TorProject.org (testnet)",
				Address: "BgqAm2xSsS2ABj8kaXc4maPkApKfWJXueN9Nw4YuAtTXbGTNykwQa7F2yCx4bGRhG1RWXoheLff6XG1JUnXtEPZFDYj1mpw"})
		donations = append(donations,
			Donation{
				Name:    "OpenPrivacy.ca (testnet)",
				Address: "Be4eYNu1XhFA3cKtF4GiPCcVmjazGtyGxDXBy6LYU5TxcGKFWBDdu8nKhfHfE4QH5kj4WtznBuuubMYJRH2MnLRbGNoNuUe",
			})
		donations = append(donations,
			Donation{
				Name:    "Dots2 (testnet)",
				Address: "Bd3FdN4gZcPYrVPyBWkboq2igocGJ673fHqrUDWmaG5h8Ypc6sQ8f8U4Hc1yV6DuoT9zTh1ekKhrG5ASsCZtzFjtNbKqRPF",
			})
		donations = append(donations,
			Donation{
				Name:    "Dots3 (testnet)",
				Address: "BftQMeTa2KB7UvMHDF6xoJGBM6xEUH9QSUuxwtiHQgDXYtA7cuj1QiW9ALKpCky8n7ZqDA5SX9rH6MxaUudeXuXCHhEufqm",
			})
		donations = append(donations,
			Donation{
				Name:    "Dots4 (testnet)",
				Address: "BanHWw8QW7pNuLgZrmTkX22XnLgPwPnN7i45jpDH5gTrM2LeRtVxKZf8v3m7Y7WrXE2RG4gvGhyY7LekU319aBCeHLULxoj",
			})
	} else {
		donations = append(donations,
			Donation{
				Name:    "TorProject.org",
				Address: "46wMHYi7ukCCe31U18DAgSbHuTRgizfxrdpRDUSuap2Abu9EiPrYMZ2ARQaH2pYHmEMX4Yd4u5VcKWaNkQf1MPXXFXq1WQc",
			})
		donations = append(donations,
			Donation{
				Name:    "OpenPrivacy.ca",
				Address: "881k3fqXxgL67djvJ5oLAJNseoWqRNBvz6oWVuLhoogCYxV6aXqtoEgHTZj7ZYmFenbAPdYDsmbjgY6AYmK66SBXPg2W7rT",
			})
		donations = append(donations,
			Donation{
				Name:    "Dots",
				Address: "88LXFRhem5PA3cKtF4GiPCcVmjazGtyGxDXBy6LYU5TxcGKFWBDdu8nKhfHfE4QH5kj4WtznBuuubMYJRH2MnLRbGPLYgTK",
			})
	}

	choices := int64(len(donations))

	myChoiceBig, err := rand.Int(rand.Reader, big.NewInt(choices))
	if err != nil {
		log.Println(err)
		myChoiceBig = big.NewInt(0)
	}
	myChoice := int(myChoiceBig.Int64())

	return donations[myChoice]
}

func GetDonationAmount(txFee uint64, txPriority walletrpc.Priority) uint64 {
	//return flatDevFee
	return txFee
}
