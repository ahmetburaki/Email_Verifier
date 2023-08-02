package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const csvFileName = "output.csv"

func main() {
	res := readFromSTDIN()
	err := writeIntoCSVFile(res)
	if err != nil {
		fmt.Printf("Error While Writing the file!\n%v", err)
	}
	fmt.Print("output.csv has been updated successfully")
}

func readFromSTDIN() [][]string {
	var res [][]string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter email addresses one by one (Enter an empty line to finish):")
	fmt.Println("CSV format will be: domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		s := checkDomain(getDomainFromEmail(text))
		res = append(res, s)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Could not read from input: %v\n", err)
	}
	return res
}

func getDomainFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func checkDomain(domain string) []string {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string = "NULL", "NULL"
	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	return []string{
		domain,
		strconv.FormatBool(hasMX),
		strconv.FormatBool(hasSPF),
		spfRecord,
		strconv.FormatBool(hasDMARC),
		dmarcRecord,
	}
}

func writeIntoCSVFile(records [][]string) error {
	var file *os.File
	var err error
	if _, err := os.Stat(csvFileName); os.IsNotExist(err) {
		// Create a new file if it doesn't exist
		file, err = os.Create(csvFileName)
		if err != nil {
			return err
		}
		w := csv.NewWriter(file)
		format := []string{"domain", "hasMX", "hasSPF", "spfRecord", "hasDMARC", "dmarcRecord"}

		// Write the header row
		err = w.Write(format)
		if err != nil {
			return err
		}

		w.Flush()

		if err := w.Error(); err != nil {
			return err
		}
	} else {
		// Open the file in append mode if it exists
		file, err = os.OpenFile(csvFileName, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}

	defer file.Close()

	// Customize the CSV writer options to include a space after each comma
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = ','    // Set the comma as the delimiter
	csvWriter.UseCRLF = true // Use CRLF line endings for better compatibility
	defer csvWriter.Flush()

	// Write the data rows
	err = csvWriter.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}
