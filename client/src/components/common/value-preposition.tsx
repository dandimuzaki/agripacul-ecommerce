"use client"

import Image from "next/image"

const ValuePreposition = () => {
  const value = [
    {image: "/grow-white.png", title: "Farm-to-Table Freshness", description: "Grown on our own land and delivered with care, every product reflects a clear origin and uncompromised quality."},
    {image: "/bowl-white.png", title: "Convenience Without Compromise", description: "Healthy food, thoughtfully prepared to fit your daily routine without sacrificing taste or nutrition."},
    {image: "/farmer-white.png", title: "Support Local Farming", description: "Each purchase contributes to a more sustainable and connected agricultural ecosystem."},
    {image: "/home-white.png", title: "Empowering Home Growers", description: "We go beyond selling food by encouraging you to grow your own and reconnect with what you consume."}
  ]
  return (
    <>
      {value.map((i) => (<div className="rounded-md flex flex-col justify-end" key={i.title}>
        <div className="h-14 w-14 md:h-20 md:w-20 flex items-center justify-center rounded-full bg-primary aspect-square mb-6">
        <Image src={i.image} alt={i.title} width={100} height={100} className="h-8 w-8 md:h-12 md:w-12 object-cover" />
        </div>
        <p className="font-bold text-lg/5 mb-2 text-primary-foreground">{i.title}</p>
        <p className="text-sm/5">{i.description}</p>
      </div>))}
    </>
  )
}

export default ValuePreposition
