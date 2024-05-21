import React, { useEffect, useState } from "react";

export default function Analytics() {
  const [url, setUrl] = useState(null);
  const [trigger, setTrigger] = useState(0);
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [msg, setMsg] = useState("");

  const [total, setTotal] = useState(0);
  const [country, setCountry] = useState(0);
  const [date, setDate] = useState(0);

  useEffect(() => {
    if (!url || !trigger) return;
    setLoading(true);
    const socket = new WebSocket(process.env.BACKEND_URL + "/analytics/" + url);

    socket.onopen = () => {
      setLoading(false);
      setSuccess(true);
      setMsg("Listening for live analytics....")
    };

    socket.onmessage = (event) => {
  
      const total = event.data.total
      const country = event.data.country
      const date = event.data.date

      setTotal(total)
    };

    socket.onclose = () => {
      console.log("WebSocket connection closed");
    };

    socket.onerror = (error) => {
      setLoading(false);
      console.error("WebSocket error:", error);
    };

    return () => {
      socket.close();
      console.log("Closed for " + url);
    };
  }, [trigger]);

  const handleChange = (e) => {
    setUrl(e.target.value);
  };

  const handleClick = (e) => {
    e.preventDefault();
    setTrigger(trigger + 1);
  };
  return (
    <div className="flex w-full justify-center items-center flex-col gap-5">
      <input
        type="text"
        placeholder="Shortened Link for analytics"
        className="border-2 border-slate-500 rounded-2xl w-full py-3 px-5"
        onChange={handleChange}
      />
      <button
        className="border-2 button w-fit  rounded-xl p-4 py-2"
        onClick={handleClick}
      >
        Generate Analytics
      </button>
      <div className={`flex w-full h-full flex-col justify-center items-center ${loading?"opacity-50":"opacity-100"}`}>
        <div className="font-semibold text-3xl">Total Vistors: 10</div>
        <div className="flex flex-wrap w-full gap-2 p-4">
          <div className="flex flex-col text-center border-2 flex-1 bg-slate-50 text-slate-900 p-5 rounded-2xl">
            <p className="w-full font-semibold text-2xl">Top 3 countries</p>
            <ul className="w-full flex flex-col justify-center">
              <li className="w-full">India</li>
              <li className="w-full">India</li>
              <li className="w-full">India</li>
            </ul>
          </div>
          <div className="flex flex-col text-center border-2 flex-1 bg-slate-50 text-slate-900 p-5 rounded-2xl">
            <p className="w-full font-semibold text-2xl">Top 3 countries</p>
            <ul className="w-full flex flex-col justify-center">
              <li className="w-full">India</li>
              <li className="w-full">India</li>
              <li className="w-full">India</li>
            </ul>
          </div>
        </div>
      </div>
      {success && <div>{msg}</div>}
    </div>
  );
}
