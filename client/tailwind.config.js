/** @type {import('tailwindcss').Config} */
export default {
    content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
    theme: {
        fontFamily: {
            FunnelSans: ['Funnel Sans', 'sans-serif'],
            Roboto: ['Roboto', 'sans-serif'],
        },
        extend: {
            colors: {
                inactive: '#CBD5E0',
                active: '#63B3ED',
                complete: '#68D391',
                incomplete: '#FC8181',
            },
        },
    },
    plugins: [],
}
