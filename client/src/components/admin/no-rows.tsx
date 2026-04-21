"use client"

export default function NoRows({colSpan, text}: {colSpan: number, text: string}) {
  return (
    <tbody>
      <tr>
        <td colSpan={colSpan} className="text-center p-4">
          No {text} match your filter
        </td>
      </tr>
    </tbody>
  )
}