"use client"

import { useEffect, useRef, useState } from "react";

const images = [
  "/pakcoy.jpg",
  "/team.jpg",
  "/garden.jpg",
  "/cherry-tomato.png",
];

export default function InfiniteSlider() {
  const [index, setIndex] = useState(0);
  const [prevIndex, setPrevIndex] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setPrevIndex(index);
      setIndex((prev) => (prev + 1) % images.length);
    }, 4000);

    return () => clearInterval(interval);
  }, [index]);

  return (
    <div className="relative w-full h-full overflow-hidden rounded-2xl">
      {images.map((img, i) => {
        let position = "translate-x-full";

        if (i === index) position = "translate-x-0";
        if (i === prevIndex) position = "-translate-x-full";

        return (
          <img
            key={i}
            src={img}
            className={`absolute w-full h-full object-cover transition-transform duration-700 ease-in-out ${position}`}
          />
        );
      })}
    </div>
  );
}