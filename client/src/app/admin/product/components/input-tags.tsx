import { Button } from '@/components/ui/button';
import { Field, FieldLabel } from '@/components/ui/field';
import { Input } from '@/components/ui/input';
import { ProductFormValues } from '@/schemas/product.schema';
import { Add, Close } from '@mui/icons-material';
import { useState } from 'react'
import { UseFormReturn } from 'react-hook-form';

const InputTags = ({ form }: { form: UseFormReturn<ProductFormValues> }) => {
  const [input, setInput] = useState("");

  const tags = form.watch("tags") ?? [];

  function addTag() {
    if (!input.trim()) return;

    form.setValue("tags", [...tags, input]);
    setInput("");
  }

  return (
    <Field>
      <FieldLabel>Tags</FieldLabel>

      <div className="flex gap-2 w-fit">
        <Input
          className='w-fit'
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              e.preventDefault();
              addTag();
            }
          }}
          placeholder="Add tag..."
        />

        <Button
          type="button"
          className="bg-yellow-500 text-white"
          onClick={addTag}
        >
          <Add fontSize="small" />
          Add Tag
        </Button>
      </div>

      <div className="flex gap-2 flex-wrap">
        {tags.map((tag, i) => (
          <div
            key={i}
            className="px-3 py-1 rounded bg-gray-200 flex items-center gap-2"
          >
            {tag}
            <button
              type="button"
              onClick={() =>
                form.setValue(
                  "tags",
                  tags.filter((_, index) => index !== i)
                )
              }
            >
              <Close fontSize="small" />
            </button>
          </div>
        ))}
      </div>
    </Field>
  );
};

export default InputTags
