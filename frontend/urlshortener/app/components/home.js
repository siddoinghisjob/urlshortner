"use client";

import React, { useEffect } from "react";
import Link from "next/link";
import Shorten from "./shorten";
import Analytics from "./analytics";
import { useSearchParams, useRouter } from "next/navigation";

export default function HomeComp() {
  const search = useSearchParams();
  const router = useRouter();
  const tab = search.get("tab");

  useEffect(() => {
    if (!tab || (tab != "shortener" && tab != "analytics")) {
      router.push("?tab=shortener", { scroll: true });
    }
  }, []);
  return (
    <div className="p-10 h-full w-full flex flex-col justify-center items-center">
      <div className="bg-white rounded-2xl w-full md:w-1/2 border-collapse p-10 flex flex-col items-center h-full border-2 justify-center">
        <div className="w-11/12 md:w-7/12 lg:w-1/3 relative">
          <Link
            href="?tab=analytics"
            className={`bg-white ${
              tab == "analytics"
                ? "-top-[5.05rem] border-b-0 z-10 cursor-not-allowed"
                : "-top-[5.2rem] border-b-1 z-0 bg-slate-50 text-slate-700 cursor-pointer"
            } absolute rounded-t-xl p-5 py-2 border-2 border-collapse left-5`}
          >
            Analytics
          </Link>
          <Link
            href="?tab=shortener"
            className={`bg-white ${
              tab != "analytics"
                ? "-top-[5.05rem] border-b-0 z-10 cursor-not-allowed"
                : "-top-[5.2rem] border-b-1 z-0 bg-slate-50 text-slate-700 cursor-pointer"
            } absolute rounded-t-xl p-5 py-2 border-2 border-collapse right-5`}
          >
            Shortener
          </Link>
        </div>
        {tab != "analytics" && <Shorten />}
        {tab == "analytics" && <Analytics />}
      </div>
    </div>
  );
}
