// pages/url/[params].js

import { headers } from "next/headers";
import { redirect } from "next/navigation";
import React from "react";

export default async function URL({ params }) {
  const country = async (ip) => {
    let msg;
    if (!ip || ip === "") {
      msg = "NA";
    }

    try {
      const res = await fetch(process.env.COUNTRY_API + "/" + ip, {
        method: "GET",
      });
      if (res.status !== 200) throw new Error();
      const json = await res.json();
      msg = json.country;
    } catch (e) {
      msg = "NA";
    }
    return msg;
  };

  const util = async (url, country) => {
    let msg;
    if (!url || url === "") {
      msg = "";
    }

    try {
      const res = await fetch(
        process.env.BACKEND_URL + "/url/" + url + "?c=" + country,
        {
          method: "GET",
          mode: "cors",
        }
      );
      if (res.status !== 200) throw new Error();
      const json = await res.json();
      if (json.error) throw new Error();
      msg = json.message;
    } catch (e) {
      msg = "";
    }
    return msg;
  };

  const header = headers();
  const ip = (header.get("x-forwarded-for") ?? "127.0.0.1").split(",")[0];

  let cName = await Promise.all([country(ip)]);
  let url = await Promise.all([util(params.url, cName[0])]);

  if (url[0] !== "") {
    let turl;
    if (url[0].startsWith("http")) turl = url[0];
    else turl = "http://" + url[0];
    redirect(turl);
  }

  return (
    <div className="flex-1 flex justify-center items-center text-3xl">
      <p>{url[0] === "" ? "Error. Not Found." : "Redirecting to " + url[0]}</p>
    </div>
  );
}
