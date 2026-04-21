"use client"

export default function Tags() {
  const tags = [
    {
      name: "vegetables",
      link: "/products?category_id=1"
    },
    {
      name: "salads",
      link: "/products?category_id=2"
    },
    {
      name: "homemade",
      link: "/products?category_id=3"
    },
    {
      name: "gardening",
      link: "/products?category_id=4"
    },
  ]

  return (
    <div className="w-full">
    <div className="flex overflow-x-auto no-scrollbar gap-2 lg:flex-wrap text-xs md:text-sm items-center">
      <p className="font-bold pr-2">Popular: </p>
      {tags.map((t) => <div key={t.name} className="cursor-pointer hover:text-primary hover:border-primary border-gray-400 px-1 py-[2px] md:px-3 md:py-1 rounded-full border">{t.name}</div>)}
    </div>
    </div>
  )
}