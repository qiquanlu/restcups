restcups
=================

The restcups is a simple REST wrapper for CUPS printing server that allows you to print binary files to a specified CUPS printer over http. 

The server listens for incoming requests and accepts binary data as input. It also accepts query parameters to specify printer options such as the number of copies and the "raw" option.

## API Endpoints
-------------
### `GET /printers`

This endpoint returns a list of available printers on the CUPS server.


#### Example Response

```
[
    {
        "name": "Brother_DCP_L2550DW_series"
    },
    {
        "name": "Brother_DCP_L2550DW_series_b42200217f10_"
    },
    {
        "name": "ZP450"
    }
]
```


### `POST /print`

This endpoint accepts binary data as input and prints it to the specified printer with the specified options. The following query parameters are supported:

#### Query Parameters
|Parameter | Type |Description| Default Value|
|----------|------|-----------|--------------|
|`printer` | String |The printer name.| required |
|`copies` | Integer |The number of copies to print.| 1 |
|`raw` | Boolean |A boolean value indicating whether to print the file in "raw" format.| false |


#### Request Body

The request body must contain binary data to print.

#### Example Request

```
`POST /print?copies=2&raw=true HTTP/1.1 Host: localhost:8080 Content-Type: application/octet-stream  [Binary Data]`
```
#### Example Response

```
`HTTP/1.1 200 OK Content-Type: application/json  {"message": "Print job submitted successfully"}`

```

## Usage
-----

### Prerequisites

*   CUPS printing system installed and running
*   firewall allow port 8080/custom port open

### Build/run from source

1.  Clone or download the Print REST Server repository from GitHub.
    
2.  Open a terminal and navigate to the root directory of the repository.
    
3.  Run the following command to start the server:
    
    
    ```
    go run server.go
    ```
    
    This will start the server on port 8080. If you want to use a different port, you can specify it using the `--port` flag. For example:
    
       ```
    go run server.go --port 8888
    ```

### Ubuntu service start on boot

1. Build or download restcups binary
    ```
    go build 
    ```

2. move to /user/bin
    ```
    sudo mv ./restcups /usr/bin
    ```

3. Setup service on boot, default port 8080, change it to different port if needed by editing restcups.service
    ```
    sudo cp ./restcups.service /etc/systemd/system && sudo chmod +x /etc/systemd/system/restcups.service
    ```
4. Start service 
    ```
    sudo systemctl start restcups
    ```

## Examples
-----

#### Printing a File

To print a file using the Print REST Server, you can use a tool such as `curl` to send a `POST` request with the file contents in the request body. For example:


```
curl -X POST -H "Content-Type: application/octet-stream" --data-binary "@label.pdf" "http://localhost:8080/print?printer=LP320&copies=2"
```

This will print the file "label.pdf" in current folder to the printer named "LP320" with 2 copies.

#### Getting Available Printers

To get a list of available printers using the Print REST Server, you can use a tool such as `curl` to send a `GET` request to the `/printers` endpoint. For example:



```
curl "http://localhost:8080/printers"
```

This will return a JSON object with a list of available printers.

License
-------

The Print REST Server is released under the [MIT License](https://opensource.org/licenses/MIT).