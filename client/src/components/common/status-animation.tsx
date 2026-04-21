"use client";

import Lottie, { LottieRefCurrentProps } from "lottie-react";
import animationData from "../../../public/loading.json";
import { useEffect, useRef } from "react";

type Status = "idle" | "loading" | "success" | "error";

export default function StatusAnimation({ status }: { status: Status }) {
  const lottieRef = useRef<LottieRefCurrentProps>(null);

  useEffect(() => {
    if (!lottieRef.current) return;

    lottieRef.current.stop();

    if (status === "loading") {
      lottieRef.current.playSegments([0, 120], true);
    }

    if (status === "success") {
      lottieRef.current.playSegments([238, 423], true);
      setTimeout(() => {
        lottieRef.current?.goToAndStop(423, true);
      }, 800);
    }

    if (status === "error") {
      lottieRef.current.playSegments([657, 841], true);
      setTimeout(() => {
        lottieRef.current?.goToAndStop(841, true);
      }, 800);
    }
  }, [status]);

  return (
    <Lottie
      lottieRef={lottieRef}
      animationData={animationData}
      loop={status === "loading"}
      className="w-40 h-40"
    />
  );
}