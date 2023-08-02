# Email_Verifier

"email_verifier" is a Go program that allows you to verify multiple email domains for the presence of essential DNS records such as MX, SPF, and DMARC. It takes email addresses as input, extracts the domain, performs DNS lookups, and stores the results in a CSV file for further analysis.

## Requirements

- Go programming language (Golang) must be installed on your system.

## How to Use

1. Clone the repository or download the source code.

2. Navigate to the project directory:

   ```
   cd email_verifier
   ```

3. Run the program:

   ```
   go run main.go
   ```
   or
   ```
   ./main
   ```

	and then
	```
	info@company.com
	 ```

4. The program will prompt you to enter email addresses one by one. Provide the email addresses and press Enter after each entry. When you're done entering email addresses, simply press Enter without typing anything to finish.

5. The program will perform DNS lookups for each domain extracted from the email addresses and display the results on the console. It will show if the domain has an MX record (mail server), SPF record (Sender Policy Framework), and DMARC record (Domain-based Message Authentication, Reporting, and Conformance).

6. The program will also store the results in a CSV file named `output.csv`. The CSV file will have the following columns:
   - domain: The email domain.
   - hasMX: Whether the domain has an MX record (true or false).
   - hasSPF: Whether the domain has an SPF record (true or false).
   - spfRecord: The SPF record associated with the domain (NULL if not found).
   - hasDMARC: Whether the domain has a DMARC record (true or false).
   - dmarcRecord: The DMARC record associated with the domain (NULL if not found).

## Note

- Ensure that you have an active internet connection to perform DNS lookups.

- The program uses the `net` package in Go for DNS lookups, which relies on the operating system's DNS resolver. Some DNS resolvers may have restrictions or caching mechanisms that can affect the results.

- The CSV file will be created or updated in the same directory as the program.

- If the CSV file (`output.csv`) already exists, the program will append the new results to it. Otherwise, it will create a new CSV file and write the results.

## Contributing

Feel free to contribute to this project by creating issues or submitting pull requests. Your suggestions and improvements are welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Credits

The program was created by [MythScapegoat | ABI](https://github.com/ahmetburaki) and is based on the [Golang](https://golang.org/) programming language.
