"use client"

import { ArrowBackIos } from "@mui/icons-material"
import Link from "next/link"

export default function BackNavigation({link}: {link: string}) {
  return (
    <Link href={link} className="flex items-center"><ArrowBackIos fontSize="inherit"/>Back</Link>
  )
}