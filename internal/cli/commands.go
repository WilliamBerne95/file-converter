package cli

import (
    "github.com/spf13/cobra"
    "github.com/yourusername/file-converter/internal/converter"
)

func Execute() {
    var rootCmd = &cobra.Command{
        Use:   "file-converter",
        Short: "Un outil de conversion de fichiers",
    }

    var csvToJSONCmd = &cobra.Command{
        Use:   "csv2json [input] [output]",
        Short: "Convertir un fichier CSV en JSON",
        Args:  cobra.ExactArgs(2),
        RunE: func(cmd *cobra.Command, args []string) error {
            input, err := os.Open(args[0])
            if err != nil {
                return err
            }
            defer input.Close()

            output, err := os.Create(args[1])
            if err != nil {
                return err
            }
            defer output.Close()

            conv := &converter.TextConverter{}
            return conv.CSVToJSON(input, output)
        },
    }

    rootCmd.AddCommand(csvToJSONCmd)
    rootCmd.Execute()
}