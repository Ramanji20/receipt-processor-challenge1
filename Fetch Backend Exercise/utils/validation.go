package utils

import (
	"regexp"
	"receipt-processor-challenge/models"
)

func ValidateReceipt(receipt models.Receipt) bool {

	// match retailer name
	if matched, _ := regexp.MatchString("^[\\w\\s\\-&]+$", receipt.Retailer); !matched {
        return false
    }

	// match date
	if matched, _ := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}$", receipt.PurchaseDate); !matched {
        return false
    }

	// match time
	if matched, _ := regexp.MatchString("^\\d{2}:\\d{2}$", receipt.PurchaseTime); !matched {
		return false;
	}

	// match total price
	if matched, _ := regexp.MatchString("^\\d+\\.\\d{2}$", receipt.Total); !matched {
		return false;
	}

	// match each item description and price
	desc_regex, _ := regexp.Compile(`^[\w\s\-]+$`)
	price_regex, _ := regexp.Compile(`^\d+\.\d{2}$`)
	for _, item := range receipt.Items {
		if matched := desc_regex.MatchString(item.ShortDescription); !matched {
			return false;
		}

		if matched := price_regex.MatchString(item.Price); !matched {
			return false;
		}
	}

	return true;
}