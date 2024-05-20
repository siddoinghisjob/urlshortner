"use client";

import React, { useEffect } from "react";

export default function URL({ params }) {
  useEffect(() => {
    setTimeout(() => {
      window.location.href="/?tab";
    }, 1000);
  }, []);
  return (
    <div>
        
    </div>
  )
}
