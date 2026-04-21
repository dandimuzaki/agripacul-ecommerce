'use client';

import { hexToRgb } from "@/lib/color";

const StockBadge = ({status}: {status: string}) => {
  function color(status: string) {
    switch (status) {
      case "low stock":
        return "#eab308"; // yellow
      case "in stock":
        return "#22c55e"; // green
      case "out of stock":
        return "#ef4444"; // red
      default:
        return "#9ca3af"; // gray
    }
  }

  const divStyle = {
    color: color(status),
    backgroundColor: `rgba( ${hexToRgb(color(status))}, 0.2 )`,
  };

  return (
    <div className={`py-1 px-3 w-fit h-fit rounded-md text-xs uppercase nowrap`}
      style={divStyle}
    >
      {status}
    </div>
  )
}

export default StockBadge
