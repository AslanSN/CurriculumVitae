# Web Curriculum Vitae **WIP**

This project is a web CV build with Go(lang), Templ, Echo and Tailwind/Ripple ui.

## Description

The aim of this project is to create a web application that serves as an online CV, where you can display your personal information, work experience, skills and projects in an attractive and easy-to-navigate way.

## Characteristics

- Modern and responsive design using Tailwind CSS
- Interactive animations to enhance the user experience
- Easy to customise and update
- Simple deployment on a web server

## Requirements

- Go versión 1.16 o superior

## Instalation

1. Clone the repo:

```bash
git clone https://github.com/AslanSN/CurriculumVitae.git
```

2. Install Go dependencies:

```bash
go get
```

4. Compile

```bash
make build
```

> For more commands go to ./Makefile, and use them: `make <command>`

## Use

1. Execute the app:

```bash
air
```

2. Open your browser and go to`http://localhost:2345` or `http://localhost:2340` if air proxy port doesn't work.

## Customization

To customize your curriculum, edit the following files:

- `views/index.templ`: Edit to change dependencies as tailwind and for example the language of your web.
- `components/*.templ`: Customise Tailwind CSS styles and add your own animations and edit HTML content according to your personal details
  > **DO NOT EDIT** `<component>_templ.go` files, they will update by their own thanks to air command.
- `main.go`, `vercel/handler.go`, `vercel.json`: Adjust application settings and routes as needed
- `db/`: There's all the info, change it to add your own info. Any change at the struct will have a long domino way impact, **be aware**.
  - `models`: There's the info adapted to use in a docker postgresql gorm database
  - `constants`: There's the info in a go *raw* state.
- `.env`: you will need an env file that matches dns configuration showed at `storage/postgres.go`

## Deploy

This web is deployed using vercel so is ready to deploy it with vercel using your own repo and link it with Vercel App.

## Contributions

If you encounter any problems or have suggestions for improvement, feel free and welcomed to open an issue or submit a pull request.

## License

This project is distributed under the [MIT License](LICENSE).
