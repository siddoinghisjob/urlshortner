import React, { useEffect, useState } from "react";

export default function Analytics() {
  const [url, setUrl] = useState(null);
  const [trigger, setTrigger] = useState(false);
  const [loading, setLoading] = useState(false);
  const [msg, setMsg] = useState(null);
  const [err, setErr] = useState(null);

  const [total, setTotal] = useState(0);
  const [country, setCountry] = useState([]);
  const [date, setDate] = useState([]);

  useEffect(() => {
    if (!url) return;

    setLoading(true);
    setMsg(null);

    let tmp = url;
    if (tmp.startsWith(process.env.FRONTEND_URL + "/")) {
      tmp = tmp.slice((process.env.FRONTEND_URL + "/").length);
    } else {
      setLoading(false);
      setErr(true);
      setMsg("Encountered Error. Wrong URL.");
      return;
    }
    const pattern = /^[a-zA-Z]{2}\d+[a-zA-Z]{2}$/;
    const isValid = pattern.test(tmp);

    if (!isValid) {
      setLoading(false);
      setErr(true);
      setMsg("Encountered Error. Wrong URL.");
      return;
    }

    const socket = new WebSocket(
      process.env.BACKEND_WS_URL + "/analytics/" + tmp
    );

    socket.onopen = () => {
      setLoading(false);
      setMsg("Listening for live analytics....");
    };

    socket.onmessage = async (event) => {
      const d = await JSON.parse(event.data);
      if (d.hasOwnProperty("error") && d.error) {
        setErr(true);
        setMsg("Not Found.");
        socket.close();
        return;
      }

      const total = d.total;
      const country = d.country;
      const date = d.date;

      setTotal(total);
      setCountry(country);
      setDate(date);
    };

    socket.onerror = (error) => {
      setLoading(false);
      setErr(true);
      setMsg("Encountered Error. 3");
    };

    return () => {
      socket.close();
      setLoading(false);
      console.log("Closed for " + url);
    };
  }, [trigger]);

  useEffect(() => console.log("Triggered"), [trigger]);

  const handleChange = (e) => {
    setUrl(e.target.value);
  };

  const handleClick = (e) => {
    e.preventDefault();
    setMsg(null);
    setTotal(0);
    setErr(false);
    setTrigger(!trigger);
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
      <div
        className={`flex relative w-full h-full flex-col justify-center items-center`}
      >
        {(err || loading) && (
          <div className="absolute top-0 bottom-0 left-0 right-0 backdrop-blur-[6px] bg-black bg-opacity-30 flex justify-center items-center text-3xl font-semibold w-full h-full rounded-3xl cur">
            {err && (
              <p className="bg-rose-50 p-2 rounded-2xl border-2 border-red-800 text-red-700">
                {msg}
              </p>
            )}
            {loading && (
              <p className="bg-rose-50 p-2 rounded-2xl border-2 border-red-800 text-red-700">
                Loading..
              </p>
            )}
          </div>
        )}
        <div className="font-semibold text-3xl">Total Vistors: {total}</div>
        <div className="flex flex-wrap w-full gap-2 p-4">
          <div className="flex flex-col text-center border-2 flex-1 bg-slate-50 text-slate-900 p-5 rounded-2xl">
            <p className="w-full font-semibold text-2xl">
              Top {country?.length} countries
            </p>
            <ul className="w-full flex flex-col">
              {country?.map((el, key) => (
                <li key={key} className="w-full">
                  <u>{el.name}</u> : {el.data}
                </li>
              ))}
            </ul>
          </div>
          <div className="flex flex-col text-center border-2 flex-1 bg-slate-50 text-slate-900 p-5 rounded-2xl">
            <p className="w-full font-semibold text-2xl">
              Top {date?.length} dates
            </p>
            <ul className="w-full flex flex-col justify-center">
              {date?.map((el, key) => (
                <li key={key} className="w-full">
                  <u>{el.name}</u> : {el.data}
                </li>
              ))}
            </ul>
          </div>
        </div>
      </div>
      {msg && !err && <div>{msg}</div>}
    </div>
  );
}
