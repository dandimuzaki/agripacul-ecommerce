"use client"

import Image from "next/image"
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "../ui/accordion"

const ValueSection = () => {
  return (
    <section>
      <Accordion type="single" collapsible>
        <AccordionItem value="item-1">
          <AccordionTrigger>Farm-to-Table Freshness</AccordionTrigger>
          <AccordionContent>
            Grown on our own land and delivered with care, every product reflects a clear origin and uncompromised quality.
          </AccordionContent>
        </AccordionItem>
        <AccordionItem value="item-2">
          <AccordionTrigger>Convenience Without Compromise</AccordionTrigger>
          <AccordionContent>
            Healthy food, thoughtfully prepared to fit your daily routine without sacrificing taste or nutrition.
          </AccordionContent>
        </AccordionItem>
        <AccordionItem value="item-3">
          <AccordionTrigger>Support Local Farming</AccordionTrigger>
          <AccordionContent>
            Each purchase contributes to a more sustainable and connected agricultural ecosystem.
          </AccordionContent>
        </AccordionItem>
        <AccordionItem value="item-4">
          <AccordionTrigger className="text-primary text-xl font-bold">
            <div className="flex gap-2 items-center">
            <Image src="/home-icon.png" alt="Fresh Vegetables by Agripacul" width={100} height={100} className="h-8 w-8 object-cover relative" />
            Empowering Home Growers
            </div>
          </AccordionTrigger>
          <AccordionContent className="">
            We go beyond selling food by encouraging you to grow your own and reconnect with what you consume.
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </section>
  )
}

export default ValueSection
