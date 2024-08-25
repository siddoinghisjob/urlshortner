import { Inter } from "next/font/google";
import "./globals.css";
import { Josefin_Sans } from "next/font/google";
import Link from "next/link";

const font = Josefin_Sans({
  weight: ["300"],
  subsets: ["latin"],
  style: ["italic"],
});
const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "URL Shortener",
  description: "Generated by create next app",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <main className=" w-full h-full min-h-screen flex flex-col justify-between">
          <h1
            className={`text-5xl mb-5 bg-slate-900 p-5 rounded-b-xl text-center w-full shadow-2xl text-slate-50 ${font.className} font-bold`}
          >
            Shortly
          </h1>
          {children}
          <div className="w-full justify-center bg-black items-center flex flex-wrap gap-3 text-white p-5 h-full border-t-2">
            Made with{" "}
            <p className="text-rose-500 font-bold text-xl">&hearts;</p>
          </div>
        </main>
      </body>
    </html>
  );
}
