from unittest.mock import patch, MagicMock, ANY
import pytest
import os

from MIB_Browser.scripts.parse_mibs import parse_mibs, FakeFileWrite


@pytest.fixture
def newFakeFileWrite():
    instance = FakeFileWrite("./")
    return instance


@pytest.mark.parametrize("file_name, mib_name", [
    ("MIB-NAME.txt", "MIB-NAME"),
    ("MIB-NAME", "MIB-NAME"),
    ("MIB_NAME", "MIB_NAME")])
def test_parse_mib(file_name, mib_name):
    with patch(
        "MIB_Browser.scripts.parse_mibs.JsonCodeGen"
    ) as mock_json_code_gen, patch(
        "MIB_Browser.scripts.parse_mibs.StubSearcher"
    ) as mock_stub_searcher, patch(
        "MIB_Browser.scripts.parse_mibs.MibCompiler"
    ) as mock_compiler, patch(
        "MIB_Browser.scripts.parse_mibs.SmiV2Parser"
    ) as mock_parser, patch(
        "MIB_Browser.scripts.parse_mibs.FakeFileWrite"
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
        mock_compiler_instance = mock_compiler.return_value

        parse_mibs(file_path, MagicMock())

        mock_stub_searcher.assert_called_once_with(*mocked_base_mibs)
        mock_file_write.assert_called_once_with(path)
        mock_file_write_instance.setOptions.assert_called_once()
        mock_compiler.assert_called_once_with(
            mock_parser_instance,
            mock_json_code_gen_instance,
            mock_file_write_instance)
        assert mock_compiler_instance.addSources.call_count == 2
        mock_compiler_instance.compile.assert_called_once_with(
            mib_name,
            noDeps=False,
            genTexts=True,
            textFilter=ANY)


def test_file_write_get_data(newFakeFileWrite):
    # TODO : write out these tests
    pass
