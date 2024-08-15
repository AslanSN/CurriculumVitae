/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./app/views/**/*.{html,js,templ,go}"],
  theme: {
    extend: {}
  },
  /** @type {import('rippleui').Config} */
  rippleui: {
    // You will have all the properties available
    themes: [{
      themeName: "dark",
      colorScheme: "dark",
      colors: {
        backgroundPrimary: "#121b22"
      }
    }]
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography')
  ]
}