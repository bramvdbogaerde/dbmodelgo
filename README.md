Go Generic Interface for Database access
======================================

This project was created to easily switch databases in projects. It's still misses some core functionality but it's still under development and actively used in real-world projects.

Currently there are two drivers available: MongoDB and RethinkDB.

## Usage

See `main.go`.

## Contributing

Steps:

1. Search for the issue and look that no one is assigned to it
2. Fork the project
3. Create a new branch with name
    * fixes-issues-ISSUENUMBER
    * feature-NAME
3. (check if all the tests succeed)
4. Create a pull request

## Roadmap

* Test all the drivers
* Add access to a generic cursor
* Generic changefeeds implementation 
* Generic inserts
* Generic removes
* Generic updates
* Create tests

## License

LGPLv2.1

With exception: library may be staticly linked to go packages
