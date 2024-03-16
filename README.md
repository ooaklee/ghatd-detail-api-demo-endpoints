<p align="center">
  <a href="https://ghatd.com">
    <img alt="GHATD" src="https://demo.ghatd.com/static/img/ghat-logo-square.png" width="180" />
  </a>
</p>
<h1 align="center">
  GHAT(D)'s Demo <code>API Detail</code> - Endpoints [WIP]
</h1>
<h4 align="center">
  Apart of the <a href="http://github.com/ooaklee/ghatd" target="_blank">GHAT(D) initiative</a>
</h4>

> This repo is strongly linked to this [**GHAT(D) PR** being merged](https://github.com/ooaklee/ghatd/pull/2)

Use GHAT(D) `Details` to hit the floor running with your next Go-base full stack web application. This `Detail` is an example endpoint that comes loaded with the demo page [**our demo page api**](https://demo.ghatd.com/api/v1/words). Use this demo to 
Kick off your project if you want to get an understanding of how you can structure the **api `Detail`** of your application


## 🚥 Getting started

`Details` are independent applications by nature that can run within the GHAT(D) framework as well as individually. To run this `api` Detail locally, please:

- Ensure you have the appropiate version of Go installed
- Run the following command:
```sh
go run api.go
```
- By default you should be able to access your `api` Detail on http://localhost:4044


For the best experience developing your Detail, we recommend using hot reloading when developing by using:
```sh
reflex -r '\.(html|go)$' -s -- go run web.go
```
You will have to ensure you have [**reflex** (click to go to installation steps)](https://github.com/cespare/reflex?tab=readme-ov-file#installation) installed.

## 🪡 Putting together your web application and the ghat(d) framework (TBC)

1.  **Create a GHAT(D) Web App.**

    Use the GHATDCLI to create a new web app, specifying this demo `api` Detail.

    ```shell
    # create a new GHAT(D) Web App using this demo api Detail
    ghatdcli new -n "my-new-web-app" -m "github.com/some-org-or-personal/my-new-web-app" -w "https://github.com/ooaklee/ghatd-detail-api-demo-endpoints"
    ```

2.  **Start developing.**

    Navigate into your new web app's directory and start it up.

    ```shell
    cd my-new-web-app/
    go mod tidy
    go run cmd/main.go start-server
    ```

    > For the best experience, we recommend using hot reloading when developing by using `reflex -r '\.(html|go)$' -s -- go run main.go start-server`. You will have to ensure you have 
    [**reflex** (click to go to installation steps)](https://github.com/cespare/reflex?tab=readme-ov-file#installation) installed.
    

3.  **Access the service's API and start editing your app code!**

    Your site is now running at `http://localhost:4000/api/v1/words`!

## License
This project is licensed under the [MIT License](./LICENSE).
