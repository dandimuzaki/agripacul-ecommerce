"use client";

import { useEffect, useState } from "react";
import { motion } from "framer-motion";
import { ArrowBack, ArrowForward } from "@mui/icons-material";

const testimonials = [
  {
    id: 1,
    name: 'Graciela & Villi',
    image: "/testimonial1.jpg",
    quote: 'The pakcoy were so crisp and fresh — it felt like they had just been picked that morning. You can really taste the difference when vegetables are grown with care.'
  },
  {
    id: 2,
    name: 'Basuki',
    image: "/testimonial2.jpg",
    quote: 'It even smelled fresh when I opened the package! Clean, vibrant, and so satisfying to eat raw.'
  },
  {
    id: 3,
    name: 'Edna, Sakura, Nada, Aimee',
    image: "/testimonial3.jpg",
    quote: 'The pakcoy had great texture — not too thick, not too soft. My mom even asked where I bought it.'
  },
  {
    id: 4,
    name: 'Diah',
    image: "/testimonial4.jpg",
    quote: 'I used it for wraps and sandwiches — it stayed crisp for days in the fridge. This is how lettuce should be.'
  }
];

export default function TestimonialCarousel() {
  const [current, setCurrent] = useState(1); // start from middle illusion
  const [isHovered, setIsHovered] = useState(false);

  // Clone edges for infinite illusion
  const extended = [
    testimonials[testimonials.length - 1],
    ...testimonials,
    testimonials[0],
  ];

  // Auto-play
  useEffect(() => {
    if (isHovered) return;

    const interval = setInterval(() => {
      setCurrent((prev) => prev + 1);
    }, 4000);

    return () => clearInterval(interval);
  }, [isHovered]);

  // Reset when reaching cloned edges
  useEffect(() => {
    if (current === extended.length - 1) {
      setTimeout(() => setCurrent(1), 300);
    }
    if (current === 0) {
      setTimeout(() => setCurrent(extended.length - 2), 300);
    }
  }, [current, extended.length]);

  const next = () => setCurrent((prev) => prev + 1);
  const prev = () => setCurrent((prev) => prev - 1);

  const CARD_WIDTH = 480;
  const GAP = 24;
  const VIEWPORT_WIDTH = 1200;

  const offset =
    current * (CARD_WIDTH + GAP) - (VIEWPORT_WIDTH / 2 - CARD_WIDTH / 2);

  return (
    <div
      className="relative w-full overflow-hidden"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      {/* GRADIENT EDGES */}
      <div className="pointer-events-none absolute inset-y-0 left-0 w-12 md:w-32 bg-gradient-to-r from-background to-transparent z-10" />
      <div className="pointer-events-none absolute inset-y-0 right-0 w-12 md:w-32 bg-gradient-to-l from-background to-transparent z-10" />

      {/* 🟢 CENTERING LAYER */}
      <div className="flex justify-center items-center">
        
        {/* 🟢 VIEWPORT (THIS IS IMPORTANT) */}
        <div className="w-[1200px] h-[240px]">
          
          {/* 🟢 YOUR CAROUSEL (motion.div stays here) */}
          <motion.div
            className="flex items-center gap-4 md:gap-6"
            animate={{ x: -offset }}
            transition={{ type: "spring", stiffness: 120, damping: 20 }}
          >
            {extended.map((item, index) => {
              const isActive = index === current;

              return (
                <div
                  key={index}
                  className={`
                    flex-shrink-0 w-[480px] transition-all duration-500
                    ${
                      isActive
                        ? "scale-110 opacity-100 z-10"
                        : "scale-90 opacity-60 blur-sm"
                    }
                  `}
                >
                  <div
                    className={`
                      bg-white rounded-lg h-32 md:h-42 mt-4 transition-all duration-500 grid grid-cols-[2fr_3fr] overflow-hidden
                      ${isActive ? "shadow-xl" : "shadow-md"}
                    `}
                  >
                    <div className="w-full h-full rounded-l-lg overflow-hidden">
                      <img src={item.image} alt={item.name} width={100} height={100} className="h-full w-full object-cover" />
                    </div>
                    <div className="flex flex-col justify-between gap-1 md:gap-2 p-2 md:p-4 h-full w-full">
                    <p className="text-gray-700 text-xs md:text-sm">{item.quote}</p>
                    <h4 className="font-semibold text-sm md:text-base">{item.name}</h4>
                    </div>
                  </div>
                </div>
              );
            })}
          </motion.div>

        </div>
      </div>

      {/* BUTTONS (stay outside viewport so they float) */}
      <button
        onClick={prev}
        className="absolute left-6 top-1/2 -translate-y-1/2 z-20 bg-white shadow-lg h-12 w-12 flex items-center justify-center rounded-full"
      >
        <ArrowBack fontSize="small"/>
      </button>

      <button
        onClick={next}
        className="absolute right-6 top-1/2 -translate-y-1/2 z-20 bg-white shadow-lg h-12 w-12 flex items-center justify-center  rounded-full"
      >
        <ArrowForward fontSize="small"/>
      </button>
    </div>
  );
}
