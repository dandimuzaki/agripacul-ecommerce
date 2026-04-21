'use client';

import { hexToRgb } from "@/lib/color";
import { capitalize } from "@/lib/formatText";

const OrderBadge = ({status}: {status: string}) => {
  function color(status: string) {
    switch (status) {
      case "waiting for payment":
        return "#f97316"; // orange
      case "processing":
        return "#eab308"; // yellow
      case "paid":
        return "#84cc16"; // lime
      case "shipped":
        return "#3b82f6"; // blue
      case "delivered":
        return "#a855f7"; // purple
      case "completed":
        return "#22c55e"; // green
      case "cancelled":
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
    <div className={`py-1 px-3 w-fit h-fit rounded-md text-xs md:text-base`}
      style={divStyle}
    >
      {capitalize(status)}
    </div>
  )
}

export default OrderBadge
