import { ModalDrawer } from "src/components/site/Modaldrawer/Modaldrawer";

import { ColourPickerField } from "@/components/ui/ColourPickerField";
import { Button } from "@/components/ui/button";
import { FormControl } from "@/components/ui/form/FormControl";
import { FormFeedback } from "@/components/ui/form/FormFeedback";
import { FormLabel } from "@/components/ui/form/FormLabel";
import { Input } from "@/components/ui/input";
import { HStack, VStack, styled } from "@/styled-system/jsx";

import { Props, useCategoryEdit } from "./useCategoryEdit";

export function CategoryEditModal(props: Props) {
  const { form, handlers } = useCategoryEdit(props);

  return (
    <ModalDrawer
      isOpen={props.isOpen}
      onClose={props.onClose}
      title="Edit category"
    >
      <styled.form
        display="flex"
        flexDir="column"
        justifyContent="space-between"
        alignItems="start"
        height="full"
        onSubmit={handlers.handleSubmit}
        gap="2"
      >
        <VStack w="full">
          <FormControl>
            <FormLabel>Name</FormLabel>
            <Input {...form.register("name")} type="text" />
            <FormFeedback error={form.formState.errors["name"]?.message}>
              The name of the category.
            </FormFeedback>
          </FormControl>

          <FormControl>
            <FormLabel>Description</FormLabel>
            <Input {...form.register("description")} type="text" />
            <FormFeedback error={form.formState.errors["description"]?.message}>
              The description for the category.
            </FormFeedback>
          </FormControl>

          <FormControl>
            <FormLabel>Colour</FormLabel>
            <ColourPickerField control={form.control} name="colour" />
            <FormFeedback error={form.formState.errors["colour"]?.message}>
              The colour for the category.
            </FormFeedback>
          </FormControl>
        </VStack>

        <HStack w="full" alignItems="center" justify="end" pb="3" gap="4">
          <Button
            type="button"
            variant="ghost"
            size="sm"
            onClick={handlers.handleCancel}
          >
            Cancel
          </Button>
          <Button type="submit" size="sm">
            Save
          </Button>
        </HStack>
      </styled.form>
    </ModalDrawer>
  );
}
