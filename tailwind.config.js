// tailwind.config.js
module.exports = {
  content: ["./app/**/*.{js,ts,jsx,tsx}", "./app/src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        wine: {
          light: "#D9A7B0",
          DEFAULT: "#8B1E3F",
          dark: "#5C1326",
        },
        grape: "#4B2E5A",
      },
    },
  },
  plugins: [],
};
