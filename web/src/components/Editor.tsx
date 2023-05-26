import { Box, Flex } from "@chakra-ui/react";
import {
  EditorComponent,
  FloatingToolbar,
  Remirror,
  TableComponents,
  TableExtension,
  ThemeProvider,
  useRemirror,
} from "@remirror/react";
import { AllStyledComponent } from "@remirror/styles/emotion";
import { PropsWithChildren } from "react";
import { ExtensionPriority } from "remirror";
import {
  BlockquoteExtension,
  BoldExtension,
  BulletListExtension,
  CodeBlockExtension,
  CodeExtension,
  HardBreakExtension,
  HeadingExtension,
  ItalicExtension,
  LinkExtension,
  ListItemExtension,
  MarkdownExtension,
  OrderedListExtension,
  StrikeExtension,
  TrailingNodeExtension,
  wysiwygPreset,
} from "remirror/extensions";

const extensions = () => [
  new LinkExtension({ autoLink: true }),
  new BoldExtension(),
  new StrikeExtension(),
  new ItalicExtension(),
  new HeadingExtension(),
  new LinkExtension(),
  new BlockquoteExtension(),
  new BulletListExtension({ enableSpine: true }),
  new OrderedListExtension(),
  new ListItemExtension({
    priority: ExtensionPriority.High,
    enableCollapsible: true,
  }),
  new CodeExtension(),
  new CodeBlockExtension({ supportedLanguages: [] }),
  new TrailingNodeExtension(),
  new TableExtension(),
  new MarkdownExtension({ copyAsMarkdown: false }),
  new HardBreakExtension(),
  ...wysiwygPreset(),
];

type Props = {
  onChange: (md: string) => void;
  value?: string;
};

export function Editor({
  children,
  onChange,
  value,
}: PropsWithChildren<Props>) {
  const { manager, state, setState, getContext } = useRemirror({
    extensions,
    stringHandler: "markdown",
    content: value,
  });

  return (
    <Box minH={32} width="full" borderRadius="2xl" mb={4}>
      <AllStyledComponent style={{ width: "100%", minHeight: "6em" }}>
        <ThemeProvider>
          <Remirror
            placeholder="Write something..."
            manager={manager}
            state={state}
            onChange={(parameter) => {
              setState(parameter.state);

              // We can't use the useHelpers hook because that can only be
              // called from a component that's a child of <Remirror>...
              const ctx = getContext();

              // This assumes the MarkdownExtension is loaded during init.
              const markdownExtension =
                ctx?.manager.getExtension(MarkdownExtension);

              const md = markdownExtension?.getMarkdown(parameter.state);

              // Note: this *looks* like a "controlled component" however it's
              // only half of it, the value from this `onChange` call cannot be
              // passed back into the component because... it's pointless!
              onChange(md as string);
            }}
          >
            <Flex flexDir="column" width="full" minHeight="6em">
              {children}

              <EditorComponent />

              <TableComponents />
              <FloatingToolbar />
            </Flex>
          </Remirror>
        </ThemeProvider>
      </AllStyledComponent>
    </Box>
  );
}
