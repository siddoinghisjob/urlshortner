import React from "react";

export default function Analytics() {
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
    </div>
  );
}
