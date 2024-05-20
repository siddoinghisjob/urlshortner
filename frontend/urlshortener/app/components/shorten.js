import React, { useState } from "react";

export default function Shorten() {
  const postURL = async () => {
    setLoader(true);
    try {
      const res = await fetch("https://urlshortner-p173.onrender.com/url", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ URL: url }),
      });
    } catch (e) {
      console.log("Error.");
    } finally {
      setLoader(false);
    }
  };

  const [url, setURL] = useState();
  const [loader, setLoader] = useState(false);
  return (
    <div className="flex w-full justify-center items-center flex-col gap-5">
      <input
        type="text"
        placeholder="Link to shorten"
        className="border-2 border-slate-500 rounded-2xl w-full py-3 px-5"
      />
      {!loader && (
        <button onClick={postURL} className="border-2 button w-fit  rounded-xl p-4 py-2">
          Shorten
        </button>
      )}
      {loader && <p>Loading....</p>}
    </div>
  );
}
