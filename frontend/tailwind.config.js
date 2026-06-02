/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{ts,tsx}'],
  theme: {
    extend: {
      colors: {
        surface: {
          DEFAULT: '#111827',
          raised: '#1f2937',
          border: '#374151',
        },
      },
    },
  },
  plugins: [],
}

