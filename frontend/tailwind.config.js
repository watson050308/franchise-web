module.exports = {
    content: [
      "./index.html",
      "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    theme: {
      extend: {
        colors: {
          'indigo': {
            500: '#6658ea',
            600: '#5646d8',
          },
          'teal': {
            400: '#3ad0af',
          },
          'pink': {
            300: '#ff6b6b',
          },
        },
        fontFamily: {
          'poppins': ['Poppins', 'sans-serif'],
        },
      },
    },
    plugins: [],
  }