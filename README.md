# BLOG REST API
## Overview
ini adalah source code rest api untuk Blog Rest Api.
untuk mencobanya tinggal clone atau download

```cli
https://github.com/wiklapandu/restapi-blog-go.git
```

## Requirements
- Golang
- Gin

## Quickstart
sebelum menggunakan ada baiknya untuk mengubah port atau host dari `API` ini.

### Change Env

<table>
    <thead>
        <tr>
            <td>
                Key ENV
            </td>
            <td>
                Description
            </td>
            <td>
                Example
            </td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>
                DB_HOST
            </td>
            <td>
                untuk mengubah host databases 
            </td>
            <td>
                DB_HOST=localhost
            </td>
        </tr>
        <tr>
            <td>
                DB_NAME
            </td>
            <td>
                untuk mengubah nama databases
            </td>
            <td>
                DB_NAME=dev_go_restapi
            </td>
        </tr>
        <tr>
            <td>
                DB_USER
            </td>
            <td>
                untuk mengubah username databases
            </td>
            <td>
                DB_USER=root
            </td>
        </tr>
        <tr>
            <td>
                DB_PASSWORD
            </td>
            <td>
                untuk mengubah password databases
            </td>
            <td>
                DB_PASSWORD=root
            </td>
        </tr>
        <tr>
            <td>
                SECRET_KEY
            </td>
            <td>
                secret key untuk JSON Web Token (JWT)
            </td>
            <td>
                SECRET_KEY="Hiram Bartoletti"
            </td>
        </tr>
    </tbody>
</table>


untuk menjalankan api ketik di cli
```cli
go run main.go
```