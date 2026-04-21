"use client";
import { motion } from 'framer-motion';

const AnimatedButton = ({children}: {children: React.ReactNode;}) => {
  return (
    <motion.button 
      whileHover={{scale: 1.05, y:-2}}
      whileTap={{scale: 0.9, y:1}}
      transition={{duration: 0.2, ease: "easeIn"}}
      className="inline-flex shrink-0 items-center justify-center gap-2 rounded-md text-sm font-medium whitespace-nowrap transition-all outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:pointer-events-none disabled:opacity-50 aria-invalid:border-destructive aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 bg-primary text-white hover:bg-primary/90 h-9 px-4 py-2 has-[>svg]:px-3"
      >
      {children}
    </motion.button>
  )
}

export default AnimatedButton
