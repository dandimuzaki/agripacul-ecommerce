import * as React from "react"
import { cn } from "@/lib/utils"
import { cva, VariantProps } from "class-variance-authority"
import { RowStatus } from "@/types/variant"

const inputVariants = cva(
  "",
  {
    variants: {
      state: {
        clean: "border-input",
        created: "border-blue-500",
        updated: "border-green-500",
        deleted: "border-red-500 opacity-50",
      },
    },
    defaultVariants: {
      state: "clean",
    },
  }
)

type InputProps =
  React.ComponentProps<"input"> &
  VariantProps<typeof inputVariants> & {
    state?: RowStatus
  }

function Input({
  className,
  type = "text",
  state = "clean",
  ...props
}: InputProps) {
  return (
    <input
      type={type}
      data-slot="input"
      className={cn(
        // variant styles
        inputVariants({ state }),

        // base styles
        "h-9 w-full min-w-0 rounded-md border bg-transparent px-3 py-1 text-base shadow-xs",
        "placeholder:text-muted-foreground",
        "selection:bg-primary selection:text-primary-foreground",
        "dark:bg-input/30",
        "outline-none transition-[color,box-shadow]",

        // file input
        "file:inline-flex file:h-7 file:border-0 file:bg-transparent file:text-sm file:font-medium",

        // disabled
        "disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50",

        // focus
        "focus-visible:ring-ring/50 focus-visible:ring-[3px]",

        // invalid state
        "aria-invalid:ring-destructive/20 aria-invalid:border-destructive dark:aria-invalid:ring-destructive/40",

        className
      )}
      {...props}
    />
  )
}

export { Input }