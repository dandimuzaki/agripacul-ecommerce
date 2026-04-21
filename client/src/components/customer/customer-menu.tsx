"use client"

import * as React from "react"
import Link from "next/link"

import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
} from "@/components/ui/navigation-menu"
import { Profile } from "@/types/profile"
import useLogout from "@/hooks/auth/useLogout"

const menu = [
  {
    href: "/orders",
    title: "My Order"
  },
  {
    href: "/profile",
    title: "My Profile"
  },
]

export function CustomerMenu({profile, color}: {profile: Profile, color: boolean}) {
  const {mutate: onLogout} = useLogout()

  return (
    <NavigationMenu>
      <NavigationMenuList>
        <NavigationMenuItem className="bg-transparent hover:bg-primary-foreground active:bg-primary-foreground rounded-lg">
          <NavigationMenuTrigger className="bg-transparent hover:bg-primary-foreground active:bg-primary-foreground rounded-lg">
            <div className="flex gap-2 items-center">
              <img src={profile?.profile_image_url ?? "/profile.png"} alt={profile?.full_name ?? "profile"} height={100} width={100} className="w-8 h-8 rounded-full"/>
              <p className={`${color ? "text-primary" : "text-white"} hover:text-white`}>Hi, {profile?.full_name?.split(' ')[0]}</p>
            </div>
          </NavigationMenuTrigger>
          <NavigationMenuContent>
            <ul className="w-96">
              {menu.map((m) => (
              <ListItem key={m.title} href={m.href} title={m.title}>
              </ListItem>
              ))}
              <li className="text-sm">
                <NavigationMenuLink onClick={() => onLogout()}>
                Logout
                </NavigationMenuLink>
              </li>
            </ul>
          </NavigationMenuContent>
        </NavigationMenuItem>
      </NavigationMenuList>
    </NavigationMenu>
  )
}

function ListItem({
  title,
  children,
  href,
  ...props
}: React.ComponentPropsWithoutRef<"li"> & { href: string }) {
  return (
    <li {...props}>
      <NavigationMenuLink asChild>
        <Link href={href}>
          <div className="flex flex-col gap-1 text-sm">
            <div className="leading-none font-medium">{title}</div>
            <div className="line-clamp-2 text-muted-foreground">{children}</div>
          </div>
        </Link>
      </NavigationMenuLink>
    </li>
  )
}
