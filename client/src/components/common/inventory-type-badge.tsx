'use client';

import { hexToRgb } from "@/lib/color";

const InventoryTypeBadge = ({type}: {type: string}) => {
  function color(type: string) {
    switch (type) {
      case "out":
        return "#eab308"; // yellow
      case "in":
        return "#22c55e"; // green
      case "adjustment":
        return "#3b82f6"; // blue
      default:
        return "#9ca3af"; // gray
    }
  }

  const divStyle = {
    color: color(type),
    backgroundColor: `rgba( ${hexToRgb(color(type))}, 0.2 )`,
  };

  return (
    <div className={`py-1 px-3 w-fit h-fit rounded-md text-sm uppercase`}
      style={divStyle}
    >
      {type}
    </div>
  )
}

export default InventoryTypeBadge
