'use client';

import { hexToRgb } from "@/lib/color";

const SKUStatusBadge = ({status}: {status: string}) => {
  function color(status: string) {
    switch (status) {
      case "inactive":
        return "#eab308"; // yellow
      case "active":
        return "#22c55e"; // green
      case "archived":
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
    <div className={`py-1 px-3 w-fit h-fit rounded-md text-xs uppercase`}
      style={divStyle}
    >
      {status}
    </div>
  )
}

export default SKUStatusBadge
