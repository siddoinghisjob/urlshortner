"use client";

import { Josefin_Sans } from "next/font/google";
import Link from "next/link";
import { useSearchParams } from "next/navigation";
import { useRouter } from "next/navigation";
import Shorten from "./shorten";
import Analytics from "./analytics";

const font = Josefin_Sans({
  weight: ["300"],
  subsets: ["latin"],
  style: ["italic"],
});

export default function Body() {
  const search = useSearchParams();
  const router = useRouter();
  const tab = search.get("tab");

  if (!tab || (tab != "shortener" && tab != "analytics")) {
    router.push("?tab=shortener", { scroll: false });
  }
  return (
    <main className=" w-full h-full min-h-screen flex flex-col justify-between">
      <h1
        className={`text-5xl bg-slate-900 p-5 rounded-b-xl text-center w-full shadow-2xl text-slate-50 ${font.className} font-bold`}
      >
        Shortly
      </h1>
      <div className="p-10 h-full w-full flex flex-col justify-center items-center">
        <div className="bg-white rounded-2xl w-full md:w-1/2 border-collapse p-10 flex flex-col items-center h-full border-2 justify-center">
          <div className="w-11/12 md:w-7/12 lg:w-1/3 relative">
            <Link
              href="?tab=analytics"
              className={`bg-white ${
                tab == "analytics"
                  ? "-top-[5.05rem] border-b-0 z-10"
                  : "-top-[5.2rem] border-b-1 z-0 bg-slate-50 text-slate-700"
              } absolute rounded-t-xl p-5 py-2 border-2 border-collapse left-5`}
            >
              Analytics
            </Link>
            <Link
              href="?tab=shortener"
              className={`bg-white ${
                tab != "analytics"
                  ? "-top-[5.05rem] border-b-0 z-10"
                  : "-top-[5.2rem] border-b-1 z-0 bg-slate-50 text-slate-700"
              } absolute rounded-t-xl p-5 py-2 border-2 border-collapse right-5`}
            >
              Shortener
            </Link>
          </div>
          {tab != "analytics" && <Shorten />}
          {tab == "analytics" && <Analytics />}
        </div>
      </div>
      <div className="w-full justify-center bg-black items-center flex flex-wrap gap-3 text-white p-5 h-full border-t-2">
        Made with <p className="text-rose-500 font-bold text-xl">&hearts;</p> by{" "}
        <Link
          href="https://github.io/siddoinghisjob"
          className="underline decoration-dotted"
        >
          Soumya Deep Sarkar
        </Link>
      </div>
    </main>
  );
}
