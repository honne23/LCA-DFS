# Lowest common ancestor using depth-first-search


Bureaucr.at is a typical hierarchical organisation. Claire, its CEO, has a hierarchy of employees reporting to her and each employee can have a list of other employees reporting to him/her. An employee with at least one report is called a Manager.

Your task is to implement a corporate directory for Bureaucr.at with an interface to find the closest common Manager (i.e. farthest from the CEO) between two employees. You may assume that all employees eventually report up to the CEO.

#### To build:
```bash
go build -o optio main.go && chmod +x optio
```

#### Run:
```bash
./optio
```
#### Or:
```bash
go run main.go
```
