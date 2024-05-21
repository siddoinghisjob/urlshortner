/** @type {import('next').NextConfig} */
const nextConfig = {
    env :{
        BACKEND_URL : "https://urlshortner-p173.onrender.com",
        BACKEND_WS_URL : "wss://urlshortner-p173.onrender.com",
        FRONTEND_URL : "https://urlshortner-mocha.vercel.app",
        COUNTRY_API : "https://api.country.is"
    }
};

export default nextConfig;
