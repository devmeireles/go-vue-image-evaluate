# Image report evaluator

This is an API that processes image report requests. Users can submit an image URL, which will be automatically analyzed by an external service responsible for classifying the image content.

- [Stack](#stack)
- [Installing](#installing)
- [Developing](#developing)
- [Executing](#executing)
- [REST API](#rest-api)
  - [Capabilities](#capabilities)
  - [Endpoints](#endpoints)
- [Ends](#ends)

## Stack

- Fiber
- GORM
- go.uuid
- Swagger
- Docker
- CompileDaemon

## Installing

Make sure to create a .env file
```
mv env.example .env
```

Build and up the API container with:

```
docker-compose up --build
```

We use Cloudmersive's image classification service to analyze images. To get started, you can create a free account on https://account.cloudmersive.com/signup. With a free tier account, you can make up to 800 API calls per month, with a maximum file size of 3.5 MB. Once you've created an account, you'll need to generate an API key and add it to the `API_KEY` property in your `.env` file

## Developing

CompileDaemon enables a seamless development process with live reloading. This means that whenever a change is made to the files in the `app` folder, the build command will be automatically executed.

Also, you can automatically write Swagger docs using annotations throughout the controllers and models, so make to sure to execute the following command:

```
docker exec -it image-evaluate-web make build-swagger
```

After that you can see the docs built on http://localhost:3000/docs/index.html


## Executing

There are some exposed endpoints allowing to execute operations on Report model

### Capabilities

<table class="demo">
  <thead>
    <tr>
      <th>Method</th>
      <th>Endpoint</th>
      <th>Capability</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><strong>GET</strong></td>
      <td>http://localhost:3000/report</td>
      <td>Returns a list of reports</td>
    </tr>
    <tr>
      <td><strong>GET</strong></td>
      <td>http://localhost:3000/report/:id</td>
      <td>Returns a report according to the id</td>
    </tr>
    <tr>
      <td><strong>POST</strong></td>
      <td>http://localhost:3000/report</td>
      <td>Creates a report</td>
    </tr>
    <tr>
      <td><strong>PUT</strong></td>
      <td>http://localhost:3000/report/:id</td>
      <td>Updates a report according to the id</td>
    </tr>
  </tbody>
</table>