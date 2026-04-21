import { Calendar } from "@/components/ui/calendar"
import { Field, FieldLabel } from "@/components/ui/field"
import { InputGroup, InputGroupAddon, InputGroupButton, InputGroupInput } from "@/components/ui/input-group"
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"
import { ProfileFormValues } from "@/schemas/profile.schema"
import { Profile } from "@/types/profile"
import { CalendarIcon } from "lucide-react"
import { useEffect, useState } from "react"
import { Controller, UseFormReturn } from "react-hook-form"

export default function BirthInput({form, onEdit, profile}: {onEdit: boolean, form: UseFormReturn<ProfileFormValues>, profile: Profile}) {
  const [open, setOpen] = useState(false)
  const [date, setDate] = useState<Date | undefined>(undefined)
  const [month, setMonth] = useState<Date | undefined>(undefined)

  const formatDate = (date?: Date) => {
    if (!date) return ""
    return date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "long",
      day: "2-digit"
    })
  }

  useEffect(() => {
    if (profile) {
      form.setValue("date_of_birth", profile.date_of_birth)
    }
  })

  return (
    <Controller
      name="date_of_birth"
      control={form.control}
      render={({ field, fieldState }) => (
        <Field data-invalid={fieldState.invalid}>
          <FieldLabel htmlFor="profile-date-of-birth">
            Date of Birth
          </FieldLabel>
          <InputGroup>
            <InputGroupInput
              disabled={!onEdit}
              id="date-required"
              value={field.value as string || ""}
              placeholder="June 01, 2025"
              onChange={(e) => {
                const date = new Date(e.target.value)
                form.setValue("date_of_birth", formatDate(date))
              }}
              onKeyDown={(e) => {
                if (e.key === "ArrowDown") {
                  e.preventDefault()
                  setOpen(true)
                }
              }}
            />
            <InputGroupAddon align="inline-end">
              <Popover open={open} onOpenChange={setOpen}>
                <PopoverTrigger asChild>
                  <InputGroupButton
                    id="date-picker"
                    variant="ghost"
                    size="icon-xs"
                    aria-label="Select date"
                  >
                    <CalendarIcon />
                    <span className="sr-only">Select date</span>
                  </InputGroupButton>
                </PopoverTrigger>
                <PopoverContent
                  className="w-auto overflow-hidden p-0"
                  align="end"
                  alignOffset={-8}
                  sideOffset={10}
                >
                  <Calendar
                    mode="single"
                    selected={field.value ? new Date(field.value as string) : undefined}
                    month={month}
                    onMonthChange={setMonth}
                    onSelect={(date) => {
                      setDate(date)
                      form.setValue("date_of_birth", formatDate(date))
                      setOpen(false)
                    }}
                  />
                </PopoverContent>
              </Popover>
            </InputGroupAddon>
          </InputGroup>
        </Field>
      )}
    />
  )
}