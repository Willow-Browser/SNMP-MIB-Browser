from unittest.mock import patch, MagicMock
import pytest
import os

from MIB_Browser.scripts.parse_mibs import parse_mibs


@pytest.mark.parametrize("file_name", [
    ("MIB-NAME.txt"),
    ("MIB-NAME"),
    ("MIB_NAME")])
def test_parse_mib(file_name):
    with patch(
        "MIB_Browser.scripts.parse_mibs.JsonCodeGen"
    ) as mock_json_code_gen, patch(
        "MIB_Browser.scripts.parse_mibs.StubSearcher"
    ) as mock_stub_searcher, patch(
        "MIB_Browser.scripts.parse_mibs.MibCompiler"
    ) as mock_compiler, patch(
        "MIB_Browser.scripts.parse_mibs.SmiV2Parser"
    ) as mock_parser, patch(
        "MIB_Browser.scripts.parse_mibs.TestFileWrite"
    ) as mock_file_write:
        mocked_base_mibs = MagicMock()

        path = os.path.dirname(os.path.realpath(__file__))
        file_path = os.path.join(path, file_name)

        mock_stub_searcher_instance = mock_stub_searcher.return_value

        mock_json_code_gen_instance = mock_json_code_gen.return_value
        mock_json_code_gen.baseMibs = mocked_base_mibs
        mock_file_write_instance = mock_file_write.return_value
        mock_file_write_instance.setOptions.return_value = \
            mock_file_write_instance
        mock_parser_instance = mock_parser.return_value

        parse_mibs(file_path, MagicMock())

        mock_stub_searcher.assert_called_once_with(*mocked_base_mibs)
        mock_file_write.assert_called_once_with(path)
        mock_file_write_instance.setOptions.assert_called_once()
        mock_compiler.assert_called_once_with(
            mock_parser_instance,
            mock_json_code_gen_instance,
            mock_file_write_instance)
