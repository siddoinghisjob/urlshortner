import { Suspense } from "react";
import HomeComp from "./components/home";


export default function Home() {

  return (
    <Suspense>
      <HomeComp/>
    </Suspense>
  );
}
