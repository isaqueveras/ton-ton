import React, { useEffect, useRef } from "react";

import { CKEditor } from "@ckeditor/ckeditor5-react";
import ClassicEditor from "@ckeditor/ckeditor5-build-classic";

interface EditorProps {
  onChange: (data: string) => void;
  value?: string;
}

export default function Editor({ onChange, value }: EditorProps) {
  return (
    <CKEditor
      editor={ClassicEditor}
      data={value}
      onChange={(event: any, editor: any) => {
        const data = editor.getData();
        onChange(data);
      }}
      config={{
        language: 'pt-BR',
        toolbar: [
          'heading',
          '|',
          'bold',
          'italic',
          'underline',
          'strikethrough',
          '|',
          'fontSize',
          'fontColor',
          'fontBackgroundColor',
          '|',
          'alignment',
          'outdent',
          'indent',
          'bulletedList',
          'numberedList',
          'blockQuote',
          '|',
          'link',
          'insertTable',
          'imageUpload',
          'mediaEmbed',
          '|',
          'undo',
          'redo',
        ],
      }}
    />
  );
}