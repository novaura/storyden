import { FormProvider } from "react-hook-form";

import { TagListField } from "@/components/thread/ThreadTagList";
import { Button } from "@/components/ui/button";
import { HStack, WStack, styled } from "@/styled-system/jsx";

import { BodyInput } from "../BodyInput/BodyInput";
import { CategorySelect } from "../CategorySelect/CategorySelect";
import { TitleInput } from "../TitleInput/TitleInput";

import { Props, useComposeForm } from "./useComposeForm";

export function ComposeForm(props: Props) {
  const { formContext, onPublish, onSave, onAssetUpload } =
    useComposeForm(props);

  return (
    <styled.form
      display="flex"
      flexDir="column"
      alignItems="start"
      w="full"
      h="full"
      gap="2"
      onSubmit={onPublish}
    >
      <FormProvider {...formContext}>
        <WStack>
          <HStack width="full">
            <CategorySelect />
            <TagListField
              name="tags"
              control={formContext.control}
              initialTags={props.initialDraft?.tags}
            />
          </HStack>

          <HStack>
            <Button
              variant="ghost"
              size="xs"
              disabled={!formContext.formState.isValid}
              onClick={onSave}
            >
              Save draft
            </Button>

            <Button
              variant="subtle"
              size="xs"
              type="submit"
              disabled={!formContext.formState.isValid}
              // isLoading={formContext.formState.isSubmitting}
            >
              Post
            </Button>
          </HStack>
        </WStack>

        <HStack width="full" justifyContent="space-between" alignItems="start">
          <HStack width="full">
            <TitleInput />
          </HStack>
        </HStack>

        <BodyInput onAssetUpload={onAssetUpload} />
      </FormProvider>
    </styled.form>
  );
}
