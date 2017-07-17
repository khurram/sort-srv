# Sort Service

Sort service will trivially sort an array of integers using your algorithm of choice.

Implemented sorting algorithms:

* [x] Insertion Sort
* [x] Selection Sort
* [ ] Bubble Sort
* [ ] Merge Sort
* [ ] Heap Sort
* [ ] Quick Sort
* [ ] Counting Sort

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```

3. Download and start the service

	```shell
	go get github.com/mistrq/sort-srv
	sort-srv
	```

	OR as a docker container

	```shell
	docker run microhq/sort-srv --registry_address=YOUR_REGISTRY_ADDRESS
	```

## The API

```shell
micro query go.micro.srv.sort Sort.InsertionSort '{"input": [5, 4, 3, 2, 1]}'
{
	"output": [
		1,
		2,
		3,
		4,
		5
	]
}
```
