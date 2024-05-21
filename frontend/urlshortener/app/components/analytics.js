import React from "react";

export default function Analytics() {
  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080");

    socket.onopen = () => {
      console.log("WebSocket connection established");
    };

    socket.onmessage = (event) => {
      console.log("Received message:", event.data);
    };

    socket.onclose = () => {
      console.log("WebSocket connection closed");
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    return () => {
      socket.close();
    };
  }, []);
  return (
    <div className="flex w-full justify-center items-center flex-col gap-5">
      <input
        type="text"
        placeholder="Shortened Link for analytics"
        className="border-2 border-slate-500 rounded-2xl w-full py-3 px-5"
      />
      <button className="border-2 button w-fit  rounded-xl p-4 py-2">
        Generate Analytics
      </button>
      <div className="flex w-full h-full flex-col justify-center items-center">
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
      <div>Listening for live analytics....</div>
    </div>
  );
}
