# SpamHammer

**Note**: This project is for the assessment of the subject CSSE6400 Software Architecture at the University of Queensland and is not intended for any other usage.

A command line interface for scanning emails for spam or malicious content.

## Installation

This package must be installed by collecting the static binary given in the releases.

## Usage

#### Help

```bash
spamhammer --help
```

```bash
spamhammer scan --help
```


#### Scanning an Email

Tickets are generated from a JSON file like so:

```json
{
  "id": "ABCD-1234",
  "email": {
    "from": "uqehugh3@uq.edu.au",
    "to": "uqehugh3@uq.edu.au",
    "subject": "CSSE6400: Cloud Assignment Help",
    "body": "Hey Valued Student\nHows the assignment going?\nRegards\nEvan Hughes",
    "headers": {
      "X-Customer-Id": "1234",
      "X-Message-Id": "ABCD-1234",
      "X-SpamHammer-Fingerprint": "0|12"
    },
    "date": "2024-01-01T12:00:00Z"
  }
}
```

The output is a path with a filename but no extension where a .json will be generated.

```bash
spamhammer scan --input examples/input.json --output examples/output
```

You can also use stdin and stdout for input and output.

For Input:

```bash
cat examples/input.json | spamhammer scan --input '-' --output examples/output
```

or 

```bash
cat examples/input.json | spamhammer scan --output examples/output
```

For Output:

```bash
spamhammer scan --input examples/input.json --output '-' > examples/output.json
```

or 

```bash
spamhammer scan --input examples/input.json > examples/output.json
```


Example Output:

```json
{
  "id": "ABCD-1234",
  "malicious": false
}
```


## Motivation

This project was created for the assessment of the subject CSSE6400 Software Architecture at the University of Queensland.  It is intended to generate an output that requires work, this version accomplishes this by computing an arbitrary BCRYPT hash which is thrown away. The program then generates a report based on the Fingerprint given to it.

## Performance Characteristics

These stats were made on a **t4g.micro** using:

todo(evan): Redo these recordings from last year.

```bash
psrecord "spamhammer ....." --log activity.txt --plot performance.png
```

| Type                | Stats                        |
|---------------------|------------------------------|
| Scan (Small Email)  | ![](performance/scan.png)  |
| Scan (medium Email) | ![](performance/scan.png)  |
| Scan (large Email)  | ![](performance/scan.png)  |

## Contributing

Contributions are welcome but the project is for the usage in an assessment so some aspects of the program are intentional to create load on the system.

## Changes

### No Releases

No releases have been made.