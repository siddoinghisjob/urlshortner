"use client";

import React, { useEffect, useState } from "react";

export default function URL({ params }) {
  const [msg, setMsg] = useState("");

  useEffect(() => {
    if (!params.url || params.url == "") {
      setMsg("Bad Request.");
      return;
    }
    const helper = async () => {
      try {
        const res = await fetch(
          process.env.BACKEND_URL + "/url/" + params.url,
          {
            method: "GET",
            mode: "cors",
          }
        );
        if (res.status != 200) throw new Error();
        const json = await res.json();
        if (json.error) throw new Error();
        setMsg("Redirecting to " + json.message);
        setTimeout(() => {
          window.location.href = json.message;
        }, 1000);
      } catch (e) {
        setMsg("Bad Request.");
      }
    };
    helper();
  }, []);
  return (
    <div className="flex-1 flex justify-center items-center text-3xl">
      {msg !== "" && <p>{msg}</p>}
      {msg === "" && <p>Finding.... </p>}
    </div>
  );
}
