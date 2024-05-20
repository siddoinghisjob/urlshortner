import React, { useState } from "react";

export default function Shorten() {
  const [response, setRes] = useState(null);

  const [url, setURL] = useState("");

  const postURL = async () => {
    setLoader(true);
    try {
      const res = await fetch(process.env.BACKEND_URL + "/url", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ URL: url }),
      });
      const jdata = await res.json();
      if (res.status != 201 || jdata?.error) throw new Error();
      setRes({
        err: false,
        msg: process.env.FRONTEND_URL + "/" + jdata.message,
      });
    } catch (e) {
      setRes({ err: true, msg: "Error. Try again." });
    } finally {
      setLoader(false);
    }
  };

  const [loader, setLoader] = useState(false);
  return (
    <div className="flex w-full justify-center items-center flex-col gap-5">
      <input
        type="text"
        placeholder="Link to shorten"
        className="border-2 border-slate-500 rounded-2xl w-full py-3 px-5"
        onChange={(e) => {
          setURL(e.target.value);
        }}
      />
      {!loader && (
        <button
          onClick={postURL}
          className="border-2 button w-fit  rounded-xl p-4 py-2"
        >
          Shorten
        </button>
      )}
      {loader && <p>Loading....</p>}
      {response && (
        <div
          className={`p-3 py-2  rounded-xl ${
            response?.err
              ? "bg-red-50 text-red-700"
              : "bg-green-50 text-green-700"
          }`}
        >
          {response?.msg}
        </div>
      )}
    </div>
  );
}
