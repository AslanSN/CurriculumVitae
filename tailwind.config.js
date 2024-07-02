/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{html,js,templ}"],
  theme: {
    extend: {}
  },
  /** @type {import('rippleui').Config} */
  rippleui: {
    // You will have all the properties available
    // themes: [{
    //   themeName: "light",
    //   colorScheme: "light",
    //   colors: {
    //     primary: "red"
    //   }
    // }]
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography')
  ]
}