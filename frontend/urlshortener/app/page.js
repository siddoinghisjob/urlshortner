import { Suspense } from "react";
import Body from "./components/body";


export default function Home() {

  return (
    <Suspense>
      <Body/>
    </Suspense>
  );
}
